package server

import (
	"context"
	"net/http"

	"github.com/subratohld/config"
)

type Server interface {
	ListenAndServe() error
	ListenAndServeTLS(certFile, keyFile string) error
	SetKeepAlivesEnabled(bool)
	Shutdown(ctx context.Context) error
}

func New(configReader config.Reader, handlers http.Handler) Server {
	return &server{
		svr: http.Server{
			Addr:    ":" + configReader.GetString("server.port"),
			Handler: handlers,
		},
	}
}

type server struct {
	svr http.Server
}

func (s *server) ListenAndServe() (err error) {
	return s.svr.ListenAndServe()
}

func (s *server) ListenAndServeTLS(certFile, keyFile string) error {
	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	return s.svr.Shutdown(ctx)
}

func (s *server) SetKeepAlivesEnabled(f bool) {
	s.svr.SetKeepAlivesEnabled(f)
}
