package service

import (
	"crypto/tls"
	"fmt"
	"httpinfo/internal/handlers"
	"httpinfo/internal/middlewares"
	"io"
	"log"
	"net"
	"net/http"
)

func CheckIpAddress(ipAddress string) error {
	parsedIp := net.ParseIP(ipAddress)
	if parsedIp == nil {
		return fmt.Errorf("IP address structure is wrong: %s", ipAddress)
	}
	return nil
}

func NewServeMux(spaPath string) *http.ServeMux {
	mux := http.NewServeMux()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	_, err := client.Get("https://golang.org/")
	if err != nil {
		fmt.Println(err)
	}

	mux.HandleFunc("/", middlewares.NewStatsMiddleware(handlers.NewHomeHandler(spaPath)))

	mux.HandleFunc("GET /api/ip/{ipAddress}", middlewares.NewStatsMiddleware(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ipAddress := r.PathValue("ipAddress")
			if err := CheckIpAddress(ipAddress); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			client := NewClient()
			response, err := client.Get(fmt.Sprintf("https://api.iplocation.net/?ip=%s", ipAddress))
			if err != nil {
				log.Printf("Request ip info error(%s): %v", ipAddress, err)
				http.Error(w, fmt.Sprintf("Can't get ip info: %s", ipAddress), http.StatusInternalServerError)
				return
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			if err != nil {
				http.Error(w, fmt.Sprintf("Can't get ip info: %s", ipAddress), http.StatusInternalServerError)
				return
			}
			w.Write(body)
			w.Header().Add("content-type", "application/json")
		})))

	return mux
}
