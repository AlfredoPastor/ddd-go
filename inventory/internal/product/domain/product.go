package domain

import (
	"context"
	"fmt"

	"github.com/AlfredoPastor/ddd-go/shared/vo"
)

type ProductRepository interface {
	Search(context.Context, vo.ID) (Product, error)
	Update(context.Context, Product) error
}

type Product struct {
	ID       vo.ID
	Name     ProductName
	Quantity ProductQuantity
}

func (p *Product) Subtract(quantityToSub ProductQuantity) error {
	p.Quantity.value = p.Quantity.value - quantityToSub.value
	if p.Quantity.value < 0 {
		return fmt.Errorf("insufficient stock")
	}
	return nil
}

func (p *Product) Add(quantityToSub ProductQuantity) error {
	p.Quantity.value = p.Quantity.value + quantityToSub.value
	if p.Quantity.value < 0 {
		return fmt.Errorf("insufficient stock")
	}
	return nil
}
