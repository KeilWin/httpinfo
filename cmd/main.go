package main

import (
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
)

const defaultAppPort = ":8080"
const defaultIndexTemplate = "web/template/index.html"

var indexTemplate *template.Template

func main() {
	var isTls bool
	var indexTemplatePath string
	flag.BoolVar(&isTls, "tls", false, "Using tls")
	flag.StringVar(&indexTemplatePath, "index-template-path", defaultIndexTemplate, "Index template path")
	flag.Parse()

	indexTemplate = template.Must(template.ParseFiles(indexTemplatePath))

	var appPort = defaultAppPort

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(HomeHandler))

	if isTls {
		log.Fatal(http.ListenAndServeTLS(appPort, "/app/server.crt", "/app/server.key", mux))
	} else {
		log.Fatal(http.ListenAndServe(appPort, mux))
	}
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
