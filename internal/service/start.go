package service

import (
	"httpinfo/internal/common"
	"httpinfo/internal/handlers"
	"httpinfo/internal/middlewares"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	serverCfg := handlers.NewServerConfig()
	common.InitCmdArgs(serverCfg)
	common.ParseCmdArgs()

	InitLogger(serverCfg)
	handlers.SetServerConfig(serverCfg)
	middlewares.LoadServerStats(serverCfg.Dump)

	mux := NewServeMux("./web/app/httpinfo/dist")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go handlers.OnShutdown(sigs)

	server := NewServer(serverCfg, mux)
	log.Fatalln(server.ListenAndServeTLS(serverCfg.Crt, serverCfg.Key))
}
