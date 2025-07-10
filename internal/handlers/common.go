package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"

	df "httpinfo/internal/defaults"
)

var serverConfig *ServerConfig
var serverStats = ServerStats{}
var indexTemplate *template.Template

func newServerStatsForDump() *ServerStatsForDump {
	return &ServerStatsForDump{
		RequestedCounter: serverStats.RequestedCounter.Load(),
	}
}

func newTemplateConfig() *TemplatesConfig {
	return &TemplatesConfig{
		App:       df.GetAppTemplatePath(),
		Header:    df.GetHeaderTemplatePath(),
		Content:   df.GetContentTemplatePath(),
		LeftSide:  df.GetLeftSideTemplatePath(),
		RightSide: df.GetRightSideTemplatePath(),
		Footer:    df.GetFooterTemplatePath(),
	}
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		TemplateCfg: newTemplateConfig(),
	}
}

func SetServerConfig(cfg *ServerConfig) {
	if _, err := os.Stat(cfg.Crt); err != nil {
		log.Fatalf("Certificate SLL file not exists: %s", cfg.Crt)
	}

	if _, err := os.Stat(cfg.Key); err != nil {
		log.Fatalf("Key SLL file not exists: %s", cfg.Key)
	}
	serverConfig = cfg
}

func LoadServerStats(dumpPath string) {
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
}

func LoadTemplates(cfg *TemplatesConfig) {
	if _, err := os.Stat(cfg.Index); err != nil {
		log.Fatalf("Template file not exists: %s", cfg.Index)
	}

	if _, err := os.Stat(cfg.App); err != nil {
		log.Fatalf("Template file not exists: %s", cfg.App)
	}
	if _, err := os.Stat(cfg.Header); err != nil {
		log.Fatalf("Template file not exists: %s", cfg.Header)
	}
	if _, err := os.Stat(cfg.Content); err != nil {
		log.Fatalf("Template file not exists: %s", cfg.Content)
	}
	if _, err := os.Stat(cfg.Footer); err != nil {
		log.Fatalf("Template file not exists: %s", cfg.Footer)
	}
	if _, err := os.Stat(cfg.LeftSide); err != nil {
		log.Fatalf("Template file not exists: %s", cfg.LeftSide)
	}
	if _, err := os.Stat(cfg.RightSide); err != nil {
		log.Fatalf("Template file not exists: %s", cfg.RightSide)
	}

	indexTemplate = template.Must(template.ParseFiles(cfg.Index, cfg.App, cfg.Header, cfg.Content, cfg.Footer, cfg.LeftSide, cfg.RightSide))
}

func OnShutdown(sigs chan os.Signal) {
	sig := <-sigs
	fmt.Printf("Received signal: %v. Performing cleanup and exiting.\n", sig)
	dumpFile, err := os.Create(serverConfig.Dump)
	if err != nil {
		log.Fatalf("Can't open: %s", serverConfig.Dump)
	}
	defer dumpFile.Close()

	serverStatsForDump := newServerStatsForDump()

	dumpJson, err := json.Marshal(serverStatsForDump)
	if err != nil {
		log.Fatalf("Can't open: %s", serverConfig.Dump)
	}
	dumpFile.Write(dumpJson)
	os.Exit(0)
}
