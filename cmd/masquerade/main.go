package main

import (
	"log/slog"
	"net/http"
	"os"
)

func main() {
	listenAddress := os.Getenv("MASQ_LISTEN_ADDRESS")
	if listenAddress == "" {
		listenAddress = "127.0.0.1:8787"
	}

	// Add mux and handlers.
	mux := http.NewServeMux()

	// `GET /healthz` handler.
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	// Everything else handler.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "path", r.URL.Path)
		w.Write([]byte("ok\n"))
	})

	slog.Info("masquerade listening", "addr", listenAddress)
	err := http.ListenAndServe(listenAddress, mux)
	if err != nil {
		slog.Error("server stopped", "err", err)
		os.Exit(1)
	}
}
