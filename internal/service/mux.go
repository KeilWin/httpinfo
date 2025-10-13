package service

import (
	"html/template"
	"httpinfo/internal/middlewares"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type RequestToSpa struct {
	IpAddress     string              `json:"ipAddress"`
	Port          string              `json:"port"`
	Protocol      string              `json:"protocol"`
	Method        string              `json:"method"`
	Host          string              `json:"host"`
	Url           string              `json:"url"`
	Headers       map[string][]string `json:"headers"`
	ContentLength string              `json:"contentLength"`
	Body          string              `json:"body"`
}

var spaTemplate *template.Template

func LoadSpaTemplate(spaPath string) {
	_, err := os.Stat(spaPath)
	if err != nil {
		log.Fatalf("Can't find spa(%s): %v", spaPath, err)
	}
	spaTemplate, err = template.ParseFiles(spaPath)
	if err != nil {
		log.Fatalf("Can't parse spa(%s): %v", spaPath, err)
	}
}

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
			defer r.Body.Close()
			limitedReader := io.LimitReader(r.Body, 1024*1024)
			buffer := make([]byte, 1024*1024)
			n, err := limitedReader.Read(buffer)
			if err != nil && err != io.EOF {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
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

			data := RequestToSpa{
				IpAddress:     hostAndPort[0],
				Port:          hostAndPort[1],
				Protocol:      r.Proto,
				Method:        r.Method,
				Host:          r.Host,
				Url:           r.URL.String(),
				Headers:       r.Header,
				ContentLength: strconv.FormatInt(r.ContentLength, 10),
				Body:          string(buffer[:n]),
			}
			spaTemplate.Execute(w, data)
		})))

	return mux
}
