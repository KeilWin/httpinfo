package main

import (
	"html/template"
	"io"
	"net/http"
)

var indexTemplate *template.Template
var serverStats = ServerStats{}

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	serverStats.RequestedCounter.Add(1)

	bodyLimitInBytes := GetHomeHandlerBodyBytesLimitInBytes()
	limitedReader := io.LimitReader(r.Body, bodyLimitInBytes)

	buffer := make([]byte, bodyLimitInBytes)
	n, err := limitedReader.Read(buffer)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	headersCountLimit := GetHomeHandlerHeadersCountLimit()
	header := make(map[string][]string, headersCountLimit)
	counter := 0
	for key, value := range r.Header {
		header[key] = value
		counter++
		if counter > int(headersCountLimit)-1 || counter > len(r.Header)-1 {
			break
		}
	}

	req := Request{
		Method:           r.Method,
		Url:              r.URL.String(),
		Proto:            r.Proto,
		Header:           header,
		Body:             string(buffer[:n]),
		ContentLength:    r.ContentLength,
		Host:             r.Host,
		RemoteAddr:       r.RemoteAddr,
		RequestURI:       r.RequestURI,
		RequestedCounter: serverStats.RequestedCounter.Load(),
	}

	indexTemplate.ExecuteTemplate(w, "index", req)

	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Sec-Fetch-Mode", "same-origin")
	w.Header().Set("Sec-Fetch-Site", "same-site")
}
