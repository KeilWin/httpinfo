package handlers

import (
	"net/http"
)

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not found", http.StatusNotFound)
}
