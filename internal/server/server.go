package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Neurasita/rest-api/internal/cfg"
	l "github.com/Neurasita/rest-api/pkg/logger"
)

// Create new server wrapper instance
func New(h http.Handler) *Server {
	addr := fmt.Sprintf("%s:%s", cfg.APP_HOST, cfg.APP_PORT)
	return &Server{
		server: &http.Server{
			Addr:    addr,
			Handler: h,
		},
	}
}

// Server wrapper command, use this server type to create app server
type Server struct {
	server *http.Server
}

// Run server on new routine
// call WaitShutdown to prevent early close
func (s *Server) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil {
			l.Error.Fatalln(err)
			panic("unable start server")
		}
	}()
	l.Info.Printf("server running on %s\n", s.server.Addr)
}

// Prevent application from early shutdown and handle graceful shutdown
func (s *Server) WaitShutdown() {
	shutdownCtx, shutdownRelease := context.
		WithTimeout(context.Background(), time.Second*10)
	defer shutdownRelease()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
	<-sigChan

	if err := s.server.Shutdown(shutdownCtx); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			l.Error.Fatalln(err)
			return
		}
	}

	l.Info.Println("server shutdown successfull...")
}
