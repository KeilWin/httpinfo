package service

import (
	"httpinfo/internal/middlewares"
	"net/http"
	"os"
	"path/filepath"
)

func NewServeMux(spaPath string) *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir(spaPath))
	mux.HandleFunc("/", middlewares.NewStatsMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := filepath.Join(spaPath, r.URL.Path)
			if info, err := os.Stat(path); err == nil && !info.IsDir() {
				fs.ServeHTTP(w, r)
				return
			}
			http.ServeFile(w, r, filepath.Join(spaPath, "index.html"))
		})))

	return mux
}
