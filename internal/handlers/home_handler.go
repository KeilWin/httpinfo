package handlers

import (
	"httpinfo/internal/defaults"
	"io"
	"net/http"
	"strings"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	serverStats.RequestedCounter.Add(1)

	bodyLimitInBytes := defaults.GetHomeHandlerBodyBytesLimitInBytes()
	limitedReader := io.LimitReader(r.Body, bodyLimitInBytes)

	buffer := make([]byte, bodyLimitInBytes)
	n, err := limitedReader.Read(buffer)
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	headersCountLimit := defaults.GetHomeHandlerHeadersCountLimit()
	header := make(map[string][]string, headersCountLimit)
	counter := 0
	for key, value := range r.Header {
		header[key] = value
		counter++
		if counter > int(headersCountLimit)-1 || counter > len(r.Header)-1 {
			break
		}
	}

	var hostAndPort []string
	if r.RemoteAddr[0] == '[' {
		hostAndPort = strings.Split(r.RemoteAddr[1:], "]:")
	} else {
		hostAndPort = strings.Split(r.RemoteAddr, ":")
	}
	if len(hostAndPort) != 2 {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := Request{
		Method:           r.Method,
		Url:              r.URL.String(),
		Proto:            r.Proto,
		Header:           header,
		Body:             string(buffer[:n]),
		ContentLength:    r.ContentLength,
		Host:             r.Host,
		RemoteAddr:       hostAndPort[0],
		RemotePort:       hostAndPort[1],
		RequestURI:       r.RequestURI,
		RequestedCounter: serverStats.RequestedCounter.Load(),
	}

	indexTemplate.ExecuteTemplate(w, "index", req)

	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Sec-Fetch-Mode", "same-origin")
	w.Header().Set("Sec-Fetch-Site", "same-site")
}
