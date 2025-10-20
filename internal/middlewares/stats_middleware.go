package middlewares

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

type ServerStats struct {
	RequestedCounter atomic.Uint64 `json:"requestedCounter"`
}

type ServerStatsForDump struct {
	RequestedCounter uint64 `json:"requestedCounter"`
}

var serverStats = ServerStats{}

func NewServerStatsForDump() *ServerStatsForDump {
	return &ServerStatsForDump{
		RequestedCounter: serverStats.RequestedCounter.Load(),
	}
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

func NewStatsMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		serverStats.RequestedCounter.Add(1)
		next.ServeHTTP(w, r)
	})
}
