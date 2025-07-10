package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	df "httpinfo/internal/defaults"
)

func main() {
	var appPort string
	var crtPath string
	var keyPath string

	var indexTemplatePath string
	appTemplatePath := df.GetAppTemplatePath()
	headerTemplatePath := df.GetHeaderTemplatePath()
	contentTemplatePath := df.GetContentTemplatePath()
	footerTemplatePath := df.GetFooterTemplatePath()
	leftSideTemplatePath := df.GetLeftSideTemplatePath()
	rigthSideTemplatePath := df.GetRightSideTemplatePath()

	var dumpPath string
	flag.StringVar(&appPort, "app-port", df.GetAppPort(), "Application port")
	flag.StringVar(&crtPath, "crt-path", df.GetCrtPath(), "Certificate SSL path")
	flag.StringVar(&keyPath, "key-path", df.GetKeyPath(), "Key SSL path")
	flag.StringVar(&indexTemplatePath, "index-template-path", df.GetIndexTemplatePath(), "Index template path")
	flag.StringVar(&dumpPath, "dump-path", df.GetDumpPath(), "Dump statistic path")
	flag.Parse()

	{
		if _, err := os.Stat(indexTemplatePath); err != nil {
			log.Fatalf("Template file not exists: %s", indexTemplatePath)
		}

		if _, err := os.Stat(appTemplatePath); err != nil {
			log.Fatalf("Template file not exists: %s", appTemplatePath)
		}
		if _, err := os.Stat(headerTemplatePath); err != nil {
			log.Fatalf("Template file not exists: %s", headerTemplatePath)
		}
		if _, err := os.Stat(contentTemplatePath); err != nil {
			log.Fatalf("Template file not exists: %s", contentTemplatePath)
		}
		if _, err := os.Stat(footerTemplatePath); err != nil {
			log.Fatalf("Template file not exists: %s", footerTemplatePath)
		}
		if _, err := os.Stat(leftSideTemplatePath); err != nil {
			log.Fatalf("Template file not exists: %s", leftSideTemplatePath)
		}
		if _, err := os.Stat(rigthSideTemplatePath); err != nil {
			log.Fatalf("Template file not exists: %s", rigthSideTemplatePath)
		}
	}

	if _, err := os.Stat(crtPath); err != nil {
		log.Fatalf("Certificate SLL file not exists: %s", crtPath)
	}

	if _, err := os.Stat(keyPath); err != nil {
		log.Fatalf("Key SLL file not exists: %s", keyPath)
	}

	if _, err := os.Stat(dumpPath); err != nil {
		serverStats.RequestedCounter.Store(0)
	} else {
		dumpFile, err := os.Open(dumpPath)
		if err != nil {
			log.Fatalf("Can't open: %s", dumpPath)
		}
		defer dumpFile.Close()

		dumpBytes, err := io.ReadAll(dumpFile)
		if err != nil {
			log.Fatalf("Can't read: %s", dumpPath)
		}

		var serverStatsForDump ServerStatsForDump

		if err = json.Unmarshal(dumpBytes, &serverStatsForDump); err != nil {
			log.Fatalf("Can't unmarshal: %s", dumpPath)
		}

		serverStats.RequestedCounter.Store(serverStatsForDump.RequestedCounter)
	}

	indexTemplate = template.Must(template.ParseFiles(indexTemplatePath, appTemplatePath, headerTemplatePath, contentTemplatePath, footerTemplatePath, leftSideTemplatePath, rigthSideTemplatePath))

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/static"))

	mux.Handle("/", http.HandlerFunc(HomeHandler))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.Handle("/favicon.ico", http.HandlerFunc(FaviconHandler))

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("Received signal: %v. Performing cleanup and exiting.\n", sig)
		dumpFile, err := os.Create(dumpPath)
		if err != nil {
			log.Fatalf("Can't open: %s", dumpPath)
		}
		defer dumpFile.Close()

		serverStatsForDump := ServerStatsForDump{
			RequestedCounter: serverStats.RequestedCounter.Load(),
		}

		dumpJson, err := json.Marshal(serverStatsForDump)
		if err != nil {
			log.Fatalf("Can't open: %s", dumpPath)
		}
		dumpFile.Write(dumpJson)
		os.Exit(0)
	}()

	log.Fatal(http.ListenAndServeTLS(appPort, crtPath, keyPath, mux))
}
