package service

import (
	"httpinfo/internal/handlers"
	"net/http"
)

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/", http.HandlerFunc(handlers.HomeHandler))
	mux.Handle("/favicon.ico", http.HandlerFunc(handlers.FaviconHandler))

	return mux
}
