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
	"syscall"
)

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
	flag.StringVar(&dumpPath, "dump-path", defaultDumpPath, "Dump statistic path")
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
	fs := http.FileServer(http.Dir("./web/static"))

	mux.Handle("/", http.HandlerFunc(HomeHandler))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
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
