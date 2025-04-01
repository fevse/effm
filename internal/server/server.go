package server

import (
	"context"
	"net"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/fevse/effm/docs"
	"github.com/fevse/effm/internal/app"
)

type Server struct {
	server *http.Server
	app    *app.EffmApp
}

func NewServer(app *app.EffmApp, host, port string) *Server {
	return &Server{
		server: &http.Server{
			Addr: net.JoinHostPort(host, port),
		},
		app: app,
	}
}

func (s *Server) Start(ctx context.Context) error {
	mux := http.NewServeMux()

	mux.Handle("GET /", s.Show())
	mux.Handle("POST /", s.Create())
	mux.Handle("DELETE /{id}", s.Delete())
	mux.Handle("PUT /{id}", s.Update())
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	s.server.Handler = mux

	s.app.Logger.Info("server started at " + s.server.Addr)

	err := s.server.ListenAndServe()
	if err != nil {
		return nil
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.app.Logger.Info("server stopped")
	return s.server.Shutdown(ctx)
}
