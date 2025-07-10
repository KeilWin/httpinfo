package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	df "httpinfo/internal/defaults"
	"httpinfo/internal/handlers"
)

func main() {
	serverCfg := handlers.ServerConfig{}
	templateCfg := handlers.TemplatesConfig{
		App:       df.GetAppTemplatePath(),
		Header:    df.GetHeaderTemplatePath(),
		Content:   df.GetContentTemplatePath(),
		LeftSide:  df.GetLeftSideTemplatePath(),
		RightSide: df.GetRightSideTemplatePath(),
		Footer:    df.GetFooterTemplatePath(),
	}

	flag.StringVar(&serverCfg.Port, "app-port", df.GetAppPort(), "Application port")
	flag.StringVar(&serverCfg.Dump, "dump-path", df.GetDumpPath(), "Dump statistic path")
	flag.StringVar(&serverCfg.Crt, "crt-path", df.GetCrtPath(), "Certificate SSL path")
	flag.StringVar(&serverCfg.Key, "key-path", df.GetKeyPath(), "Key SSL path")
	flag.StringVar(&templateCfg.Index, "index-template-path", df.GetIndexTemplatePath(), "Index template path")
	flag.Parse()

	if _, err := os.Stat(serverCfg.Crt); err != nil {
		log.Fatalf("Certificate SLL file not exists: %s", serverCfg.Crt)
	}

	if _, err := os.Stat(serverCfg.Key); err != nil {
		log.Fatalf("Key SLL file not exists: %s", serverCfg.Key)
	}

	handlers.SetServerConfig(&serverCfg)
	handlers.LoadServerStats(serverCfg.Dump)
	handlers.LoadTemplates(&templateCfg)

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
