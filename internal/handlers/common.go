package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"httpinfo/internal/middlewares"
)

var serverConfig *ServerConfig

func NewServerConfig() *ServerConfig {
	return &ServerConfig{}
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

func OnShutdown(sigs chan os.Signal) {
	sig := <-sigs
	fmt.Printf("Received signal: %v. Performing cleanup and exiting.\n", sig)
	dumpFile, err := os.Create(serverConfig.Dump)
	if err != nil {
		log.Fatalf("Can't open: %s", serverConfig.Dump)
	}
	defer dumpFile.Close()

	serverStatsForDump := middlewares.NewServerStatsForDump()

	dumpJson, err := json.Marshal(serverStatsForDump)
	if err != nil {
		log.Fatalf("Can't open: %s", serverConfig.Dump)
	}
	dumpFile.Write(dumpJson)
	os.Exit(0)
}
