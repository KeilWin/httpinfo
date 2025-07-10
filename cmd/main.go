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

	log.Fatal(http.ListenAndServeTLS(serverCfg.Port, serverCfg.Crt, serverCfg.Key, mux))
}
