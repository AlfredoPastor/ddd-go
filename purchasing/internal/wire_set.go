//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/config"
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/application/creator"
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/infraestructure/bus"
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/infraestructure/httpController"
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/infraestructure/store"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/ginhttp"
	"github.com/AlfredoPastor/ddd-go/shared/mongoatlas"
	"github.com/AlfredoPastor/ddd-go/shared/rabbitmq"
	"github.com/google/wire"
)

var superSet = wire.NewSet(
	httpController.NewHttpController,
	bus.NewBus,
	creator.NewOrderCreatorService,
	store.NewOrderRepository,
	config.NewConfig,
	rabbitmq.NewRabbitClient,
	ginhttp.NewHttpServer,
	mongoatlas.NewMongoClient,
	NewServer,
	wire.Bind(new(eventbus.Bus), new(bus.Bus)),
	wire.Bind(new(rabbitmq.Config), new(config.Config)),
	wire.Bind(new(ginhttp.Config), new(config.Config)),
	wire.Bind(new(mongoatlas.Config), new(config.Config)),
	wire.Bind(new(domain.OrderRepository), new(store.OrderRepository)),
)
