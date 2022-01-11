package internal

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/infraestructure/bus"
	"github.com/AlfredoPastor/ddd-go/shared/ginhttp"
)

type Server struct {
	ginhttp.HttpServer
	bus.Bus
}

func NewServer(
	httpserver ginhttp.HttpServer,
	bus bus.Bus,
) Server {
	return Server{
		HttpServer: httpserver,
		Bus:        bus,
	}
}

func (s *Server) Run() error {
	ctx := context.Background()
	err := s.Bus.Run(ctx)
	if err != nil {
		return err
	}

	return s.HttpServer.Run(ctx)
}
