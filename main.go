package main

import (
	"bytes"
	"context"
	"flag"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "Path to the configuration file")
}

func main() {
	flag.Parse()

	config, err := loadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

	s := &Server{
		config: config,
		logger: logger,
	}
	http.HandleFunc("/", s.handle)

	logger.Info("Starting server", "addr", config.Addr)
	log.Fatal(http.ListenAndServe(config.Addr, nil))
}

type Server struct {
	config *Config
	logger *slog.Logger
}

func (s *Server) handle(w http.ResponseWriter, r *http.Request) {
	// Copy request body
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Try providers in order
	for _, provider := range s.config.Providers {
		err := s.tryProvider(provider, w, r, &buf)
		if err == nil {
			return
		}
		s.logger.Error("Provider failed", "provider", provider.Name, "error", err)
	}

	http.Error(w, "All providers failed", http.StatusServiceUnavailable)
}

func (s *Server) tryProvider(provider ProviderConfig, w http.ResponseWriter, r *http.Request, buf *bytes.Buffer) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(provider.TimeoutMs)*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, r.Method, provider.BaseURL+r.URL.Path, bytes.NewReader(buf.Bytes()))
	if err != nil {
		return err
	}
	s.logger.Debug("Request", "method", req.Method, "url", req.URL.String())

	// Copy headers
	for key, values := range r.Header {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	// Set API key
	req.Header.Set("Authorization", "Bearer "+provider.APIKey)

	// Make request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Copy response
	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	return err
}
