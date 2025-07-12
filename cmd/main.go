package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"httpinfo/internal/common"
	"httpinfo/internal/handlers"
)

func main() {
	serverCfg := handlers.NewServerConfig()
	common.InitCmdArgs(serverCfg)
	common.ParseCmdArgs()

	logFile, err := os.OpenFile(serverCfg.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile | log.LUTC | log.Lmsgprefix)
	handlers.SetServerConfig(serverCfg)
	handlers.LoadServerStats(serverCfg.Dump)
	handlers.LoadTemplates(serverCfg.TemplateCfg)

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.Handle("/", http.HandlerFunc(handlers.HomeHandler))
	mux.Handle("/favicon.ico", http.HandlerFunc(handlers.FaviconHandler))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go handlers.OnShutdown(sigs)

	server := &http.Server{
		Addr:    serverCfg.Port,
		Handler: mux,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,

			SessionTicketsDisabled: false,
			Renegotiation:          tls.RenegotiateNever,

			NextProtos: []string{"h2", "http/1.1"},
		},
		MaxHeaderBytes: 1 << 18,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
	}

	log.Fatalln(server.ListenAndServeTLS(serverCfg.Crt, serverCfg.Key))
}
