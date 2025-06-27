package main

import (
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

const defaultAppPort = ":8080"
const defaultIndexTemplate = "web/template/index.html"
const defaultCrtPath = "/app/server.crt"
const defaultKeyPath = "/app/server.key"

var indexTemplate *template.Template

func main() {
	var crtPath string
	var keyPath string
	var indexTemplatePath string
	var appPort string
	flag.StringVar(&appPort, "app-port", defaultAppPort, "Application port")
	flag.StringVar(&crtPath, "crt-path", defaultCrtPath, "Certificate SSL path")
	flag.StringVar(&keyPath, "key-path", defaultKeyPath, "Key SSL path")
	flag.StringVar(&indexTemplatePath, "index-template-path", defaultIndexTemplate, "Index template path")
	flag.Parse()

	if _, err := os.Stat(indexTemplatePath); err != nil {
		log.Fatalf("Template file not exists: %s", indexTemplatePath)
	}

	indexTemplate = template.Must(template.ParseFiles(indexTemplatePath))

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(HomeHandler))

	log.Fatal(http.ListenAndServeTLS(appPort, crtPath, keyPath, mux))
}

type Request struct {
	Method        string
	Url           string
	Proto         string
	Header        map[string][]string
	Body          string
	ContentLength int64
	Host          string
	RemoteAddr    string
	RequestURI    string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

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
		Method:        r.Method,
		Url:           r.URL.String(),
		Proto:         r.Proto,
		Header:        r.Header,
		Body:          string(buffer[:n]),
		ContentLength: r.ContentLength,
		Host:          r.Host,
		RemoteAddr:    r.RemoteAddr,
		RequestURI:    r.RequestURI,
	}

	indexTemplate.Execute(w, req)

	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Sec-Fetch-Mode", "same-origin")
	w.Header().Set("Sec-Fetch-Site", "same-site")
}
