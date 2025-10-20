package service

import (
	"httpinfo/internal/handlers"
	"httpinfo/internal/middlewares"
	"net/http"
)

func NewServeMux(spaPath string) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", middlewares.NewStatsMiddleware(handlers.NewHomeHandler(spaPath)))
	mux.HandleFunc("GET /api/ip/{ipAddress}", middlewares.NewStatsMiddleware(handlers.NewIpHandler(NewClient)))

	return mux
}
