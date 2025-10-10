package service

import (
	"httpinfo/internal/common"
	"httpinfo/internal/handlers"
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
	handlers.LoadServerStats(serverCfg.Dump)
	handlers.LoadTemplates(serverCfg.TemplateCfg)

	mux := NewServeMux()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go handlers.OnShutdown(sigs)

	server := NewServer(serverCfg, mux)
	log.Fatalln(server.ListenAndServeTLS(serverCfg.Crt, serverCfg.Key))
}
