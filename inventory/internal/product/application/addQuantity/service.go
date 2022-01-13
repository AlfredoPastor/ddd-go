package addQuantity

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/inventory/internal/product/domain"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
)

type AddQuantityService struct {
	domain.ProductRepository
	eventbus.Bus
}

func (a AddQuantityService) Add(ctx context.Context, productID vo.ID, quantity domain.ProductQuantity) error {
	product, err := a.ProductRepository.Search(ctx, productID)
	if err != nil {
		return err
	}
	product.Add(quantity)
	err = a.ProductRepository.Update(ctx, product)
	if err != nil {
		return err
	}
	err = a.Bus.Publish(ctx, []eventbus.Event{})
	if err != nil {
		return err
	}

	return nil
}

func NewAddQuantityService(repo domain.ProductRepository, bus eventbus.Bus) AddQuantityService {
	return AddQuantityService{
		ProductRepository: repo,
		Bus:               bus,
	}
}
