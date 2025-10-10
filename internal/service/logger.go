package service

import (
	"httpinfo/internal/handlers"
	"log"
	"os"
)

func InitLogger(serverCfg *handlers.ServerConfig) {
	logFile, err := os.OpenFile(serverCfg.Log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile | log.LUTC | log.Lmsgprefix)
}
