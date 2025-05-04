package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fromjyce/Nublink/internal/storage"
	"github.com/fromjyce/Nublink/pkg/config"
)

type Server struct {
	cfg       *config.Config
	fileStore *storage.FileStorage
	server    *http.Server
}

func NewServer(cfg config.Config, fileStore *storage.FileStorage) *Server {
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	server := &Server{
		cfg:       &cfg,
		fileStore: fileStore,
		server:    srv,
	}

	mux.HandleFunc("/download/", server.downloadHandler)
	return server
}

func (s *Server) Start() {
	log.Printf("Starting server on :%d\n", s.cfg.Port)
	if err := s.server.ListenAndServeTLS(s.cfg.TLSCert, s.cfg.TLSKey); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v\n", err)
	}
}
