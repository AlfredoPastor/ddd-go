package creator

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
)

type OrderCreatorService struct {
	domain.OrderRepository
	eventbus.Bus
}

func (c OrderCreatorService) Create(ctx context.Context, id, clientId, address string, orderLines []domain.OrderLine) error {
	order, err := domain.NewOrder(id, clientId, address, orderLines)
	if err != nil {
		return err
	}
	err = c.OrderRepository.Save(ctx, order)
	if err != nil {
		return err
	}
	err = c.Bus.Publish(ctx, []eventbus.Event{domain.NewOrderCreatedEvent(order)})
	if err != nil {
		return err
	}
	return nil
}

func NewOrderCreatorService(repo domain.OrderRepository, bus eventbus.Bus) OrderCreatorService {
	return OrderCreatorService{
		OrderRepository: repo,
		Bus:             bus,
	}
}
