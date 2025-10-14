package service

import (
	"crypto/tls"
	"httpinfo/internal/handlers"
	"net/http"
	"time"
)

func NewTlsConfig() *tls.Config {
	return &tls.Config{
		MinVersion: tls.VersionTLS12,
		MaxVersion: tls.VersionTLS13,
		CipherSuites: []uint16{
			// Safe ciphers from https://www.ssllabs.com/
			// TLS 1.2
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			// TLS 1.3
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
		},
		CurvePreferences: []tls.CurveID{
			tls.CurveP521,
			tls.X25519,
		},

		SessionTicketsDisabled: false,
		Renegotiation:          tls.RenegotiateNever,

		NextProtos: []string{"h2", "http/1.1"},
	}
}

func NewTransport() *http.Transport {
	return &http.Transport{
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
			MaxVersion: tls.VersionTLS13,
			CipherSuites: []uint16{
				// Safe ciphers from https://www.ssllabs.com/
				// TLS 1.2
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				// TLS 1.3
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_CHACHA20_POLY1305_SHA256,
			},
			CurvePreferences: []tls.CurveID{
				tls.CurveP521,
				tls.X25519,
			},

			SessionTicketsDisabled: false,
			Renegotiation:          tls.RenegotiateNever,

			NextProtos: []string{"http/1.1"},
		},
	}
}

func NewServer(serverCfg *handlers.ServerConfig, mux *http.ServeMux) *http.Server {
	return &http.Server{
		Addr:           serverCfg.Port,
		Handler:        mux,
		TLSConfig:      NewTlsConfig(),
		MaxHeaderBytes: 1 << 18,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    10 * time.Second,
	}
}

func NewClient() *http.Client {
	return &http.Client{
		Transport: NewTransport(),
	}
}
