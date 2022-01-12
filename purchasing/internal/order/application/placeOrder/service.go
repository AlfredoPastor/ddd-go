package creator

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
)

type PlaceOrderService struct {
	domain.OrderRepository
	eventbus.Bus
}

func (p PlaceOrderService) Do(ctx context.Context, id vo.ID) error {
	order, err := p.OrderRepository.Search(ctx, id)
	if err != nil {
		return err
	}
	order.State.ChangeToPlaced()
	err = domain.BookProducts(ctx, p.OrderRepository, p.Bus, order.OderLines)
	if err != nil {
		return err
	}
	err = p.OrderRepository.Save(ctx, order)
	if err != nil {
		return err
	}
	err = p.Bus.Publish(ctx, []eventbus.Event{domain.NewOrderPlacedEvent(order)})
	if err != nil {
		return err
	}

	return nil
}

func NewPlaceOrderService(repo domain.OrderRepository, bus eventbus.Bus) PlaceOrderService {
	return PlaceOrderService{
		OrderRepository: repo,
		Bus:             bus,
	}
}
