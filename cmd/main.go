package main

import (
	"log"
	"net/http"
)

const defaultAppPort = ":8080"

func main() {
	var appPort = defaultAppPort

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(HomeHandler))

	log.Fatal(http.ListenAndServe(appPort, mux))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Connection", "close")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Sec-Fetch-Mode", "same-origin")
	w.Header().Set("Sec-Fetch-Site", "same-site")
}
