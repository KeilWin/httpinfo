package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	log.Fatalln(http.ListenAndServeTLS(serverCfg.Port, serverCfg.Crt, serverCfg.Key, mux))
}
