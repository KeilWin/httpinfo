package handlers

import (
	"sync/atomic"
)

type ServerConfig struct {
	Port string
	Dump string
	// SSL
	Crt string
	Key string

	TemplateCfg *TemplatesConfig
}

type ServerStats struct {
	RequestedCounter atomic.Uint64 `json:"requestedCounter"`
}

type ServerStatsForDump struct {
	RequestedCounter uint64 `json:"requestedCounter"`
}

type TemplatesConfig struct {
	Index string

	App string

	Header string

	Content   string
	LeftSide  string
	RightSide string

	Footer string
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
