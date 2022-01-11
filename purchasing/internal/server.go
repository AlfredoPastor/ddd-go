package internal

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/infraestructure/bus"
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/infraestructure/httpController"
)

type Server struct {
	httpController.HttpController
	bus.Bus
}

func NewServer(
	httpserver httpController.HttpController,
	bus bus.Bus,
) Server {
	return Server{
		HttpController: httpserver,
		Bus:            bus,
	}
}

func (s *Server) Run() error {
	ctx := context.Background()
	err := s.Bus.Run(ctx)
	if err != nil {
		return err
	}

	return s.HttpController.Run(ctx)
}
