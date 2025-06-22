package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

const defaultAppPort = ":8080"
const defaultIndexTemplate = "web/template/index.html"

var indexTemplate = template.Must(template.ParseFiles(defaultIndexTemplate))

func main() {
	var appPort = defaultAppPort

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(HomeHandler))

	log.Fatal(http.ListenAndServe(appPort, mux))
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

	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req := Request{
		Method:        r.Method,
		Url:           r.URL.String(),
		Proto:         r.Proto,
		Header:        r.Header,
		Body:          string(data),
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
