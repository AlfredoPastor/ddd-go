package internal

import (
	"context"
	"os"
	"os/signal"

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

func (s *Server) Run(ctx context.Context) error {
	ctx = serverContext(ctx)
	err := s.Bus.Run(ctx)
	if err != nil {
		return err
	}

	return s.HttpController.Run(ctx)
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()
	return ctx
}
