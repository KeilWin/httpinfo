package main

import "sync/atomic"

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
