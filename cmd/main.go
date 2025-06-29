package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

const defaultAppPort = ":8080"
const defaultIndexTemplate = "web/template/index.html"
const defaultCrtPath = "/etc/ssl/server.crt"
const defaultKeyPath = "/etc/ssl/server.key"
const defaultDumpPath = "/app/stats/dump.json"

var indexTemplate *template.Template
var serverStats = ServerStats{}

func main() {
	var appPort string
	var crtPath string
	var keyPath string
	var indexTemplatePath string
	var dumpPath string
	flag.StringVar(&appPort, "app-port", defaultAppPort, "Application port")
	flag.StringVar(&crtPath, "crt-path", defaultCrtPath, "Certificate SSL path")
	flag.StringVar(&keyPath, "key-path", defaultKeyPath, "Key SSL path")
	flag.StringVar(&indexTemplatePath, "index-template-path", defaultIndexTemplate, "Index template path")
	flag.StringVar(&dumpPath, "dump-path", defaultDumpPath, "Dump path")
	flag.Parse()

	if _, err := os.Stat(indexTemplatePath); err != nil {
		log.Fatalf("Template file not exists: %s", indexTemplatePath)
	}

	if _, err := os.Stat(crtPath); err != nil {
		log.Fatalf("Certificate SLL file not exists: %s", crtPath)
	}

	if _, err := os.Stat(keyPath); err != nil {
		log.Fatalf("Key SLL file not exists: %s", keyPath)
	}

	if _, err := os.Stat(dumpPath); err != nil {
		serverStats.RequestedCounter.Store(0)
	} else {
		dumpFile, err := os.Open(dumpPath)
		if err != nil {
			log.Fatalf("Can't open: %s", dumpPath)
		}
		defer dumpFile.Close()

		dumpBytes, err := io.ReadAll(dumpFile)
		if err != nil {
			log.Fatalf("Can't read: %s", dumpPath)
		}

		var serverStatsForDump ServerStatsForDump

		if err = json.Unmarshal(dumpBytes, &serverStatsForDump); err != nil {
			log.Fatalf("Can't unmarshal: %s", dumpPath)
		}

		serverStats.RequestedCounter.Store(serverStatsForDump.RequestedCounter)
	}

	indexTemplate = template.Must(template.ParseFiles(indexTemplatePath))

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(HomeHandler))
	mux.Handle("/favicon.ico", http.HandlerFunc(FaviconHandler))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %v. Performing cleanup and exiting.\n", sig)
		dumpFile, err := os.Create(dumpPath)
		if err != nil {
			log.Fatalf("Can't open: %s", dumpPath)
		}
		defer dumpFile.Close()

		serverStatsForDump := ServerStatsForDump{
			RequestedCounter: serverStats.RequestedCounter.Load(),
		}

		dumpJson, err := json.Marshal(serverStatsForDump)
		if err != nil {
			log.Fatalf("Can't open: %s", dumpPath)
		}
		dumpFile.Write(dumpJson)
		os.Exit(0)
	}()

	log.Fatal(http.ListenAndServeTLS(appPort, crtPath, keyPath, mux))
}

type ServerStats struct {
	RequestedCounter atomic.Uint64 `json:"requestedCounter"`
}

type ServerStatsForDump struct {
	RequestedCounter uint64 `json:"requestedCounter"`
}

type Request struct {
	Method           string
	Url              string
	Proto            string
	Header           map[string][]string
	Body             string
	ContentLength    int64
	Host             string
	RemoteAddr       string
	RequestURI       string
	RequestedCounter uint64
}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	serverStats.RequestedCounter.Add(1)

	limitedReader := io.LimitReader(r.Body, 1000)

	buffer := make([]byte, 1000)
	n, err := limitedReader.Read(buffer)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	header := make(map[string][]string, 50)
	counter := 0
	for key, value := range r.Header {
		header[key] = value
		counter++
		if counter > len(header)-1 {
			break
		}
	}

	req := Request{
		Method:           r.Method,
		Url:              r.URL.String(),
		Proto:            r.Proto,
		Header:           r.Header,
		Body:             string(buffer[:n]),
		ContentLength:    r.ContentLength,
		Host:             r.Host,
		RemoteAddr:       r.RemoteAddr,
		RequestURI:       r.RequestURI,
		RequestedCounter: serverStats.RequestedCounter.Load(),
	}

	indexTemplate.Execute(w, req)

	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Sec-Fetch-Mode", "same-origin")
	w.Header().Set("Sec-Fetch-Site", "same-site")
}
