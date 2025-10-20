package service

import (
	"fmt"
	"httpinfo/internal/common"
	df "httpinfo/internal/defaults"
	"httpinfo/internal/handlers"
	"httpinfo/internal/middlewares"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func GetSpaTemplatePath(spaDir string) string {
	return fmt.Sprintf("%s/%s", spaDir, df.GetSpaTemplateFile())
}

func Start() {
	serverCfg := handlers.NewServerConfig()
	common.InitCmdArgs(serverCfg)
	common.ParseCmdArgs()

	InitLogger(serverCfg)
	handlers.SetServerConfig(serverCfg)
	handlers.LoadSpaTemplate(GetSpaTemplatePath(serverCfg.Spa))
	middlewares.LoadServerStats(serverCfg.Dump)

	mux := NewServeMux(serverCfg.Spa)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go handlers.OnShutdown(sigs)

	server := NewServer(serverCfg, mux)
	log.Fatalln(server.ListenAndServeTLS(serverCfg.Crt, serverCfg.Key))
}
