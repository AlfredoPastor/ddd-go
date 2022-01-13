package substractQuantity

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/inventory/internal/product/domain"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
)

type SubstractQuantityService struct {
	domain.ProductRepository
	eventbus.Bus
}

func (s SubstractQuantityService) Subtract(ctx context.Context, productID vo.ID, quantity domain.ProductQuantity) error {
	product, err := s.ProductRepository.Search(ctx, productID)
	if err != nil {
		return err
	}
	product.Subtract(quantity)
	err = s.ProductRepository.Update(ctx, product)
	if err != nil {
		return err
	}
	err = s.Bus.Publish(ctx, []eventbus.Event{})
	if err != nil {
		return err
	}

	return nil
}

func NewSubstractQuantityService(repo domain.ProductRepository, bus eventbus.Bus) SubstractQuantityService {
	return SubstractQuantityService{
		ProductRepository: repo,
		Bus:               bus,
	}
}
