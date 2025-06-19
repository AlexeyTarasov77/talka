// Package httpserver implements HTTP server.
package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	_defaultAddr            = ":80"
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 5 * time.Second
	_defaultShutdownTimeout = 3 * time.Second
)

// Server -.
type Server struct {
	App    *gin.Engine
	server *http.Server
	notify chan error

	address         string
	prefork         bool
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeout time.Duration
}

// New -.
func New(opts ...Option) *Server {
	s := &Server{
		App:             nil,
		notify:          make(chan error, 1),
		address:         _defaultAddr,
		readTimeout:     _defaultReadTimeout,
		writeTimeout:    _defaultWriteTimeout,
		shutdownTimeout: _defaultShutdownTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(s)
	}

	s.App = gin.Default()

	s.server = &http.Server{
		Addr:         s.address,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
		Handler:      s.App.Handler(),
	}

	return s
}

// Start -.
func (s *Server) Start() {
	go func() {
		s.notify <- s.App.Run(s.address)
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify -.
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown -.
func (s *Server) Shutdown() error {
	ctx, _ := context.WithTimeout(context.Background(), s.shutdownTimeout)
	return s.server.Shutdown(ctx)
}
