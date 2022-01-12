package domain

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/shared/vo"
)

const TAX_PERCENT float64 = 19.0

type OrderRepository interface {
	SearchByClient(context.Context, vo.ID) (Order, error)
	Search(context.Context, vo.ID) (Order, error)
	Delete(context.Context, vo.ID) error
	Save(context.Context, Order) error
}

type Order struct {
	ID              vo.ID
	ClientID        vo.ID
	AddressShipping OrderAddressShipping
	Taxes           OrderTaxes
	Subtotal        OrderSubtotal
	Total           OrderTotal
	OderLines       []OrderLine
}

func NewOrder(id, clientId, address string, orderLines []OrderLine) (Order, error) {
	idVo, err := vo.NewIDFromString(id)
	if err != nil {
		return Order{}, err
	}
	idClientVo, err := vo.NewIDFromString(clientId)
	if err != nil {
		return Order{}, err
	}
	addressVo, err := NewOrderAddressShipping(address)
	if err != nil {
		return Order{}, err
	}
	order := Order{
		ID:              idVo,
		ClientID:        idClientVo,
		AddressShipping: addressVo,
		OderLines:       orderLines,
	}
	err = order.MakeCalculation()
	if err != nil {
		return Order{}, err
	}

	return order, nil
}

func (o *Order) MakeCalculation() error {
	sumTaxes := 0.0
	sumSubtotal := 0.0
	for _, ordLine := range o.OderLines {
		subtotal := float64(ordLine.Quantity.Primitive()) * ordLine.Price.Primitive()
		sumTaxes = sumTaxes + (subtotal / TAX_PERCENT)
		sumSubtotal = sumSubtotal + subtotal
	}
	taxesVo, err := NewOrderTaxes(sumTaxes)
	if err != nil {
		return err
	}
	subtotalVo, err := NewOrderSubtotal(sumSubtotal)
	if err != nil {
		return err
	}
	totalVo, err := NewOrderTotal(sumSubtotal + sumTaxes)
	if err != nil {
		return err
	}
	o.Taxes = taxesVo
	o.Subtotal = subtotalVo
	o.Total = totalVo

	return nil
}

type OrderLine struct {
	ID        vo.ID
	ProductID vo.ID
	Price     OrderLinePrice
	Quantity  OrderLineQuantity
}

func NewOrderLine(id, productId string, price float64, quantity int) (OrderLine, error) {
	idVo, err := vo.NewIDFromString(id)
	if err != nil {
		return OrderLine{}, err
	}
	productIdVo, err := vo.NewIDFromString(productId)
	if err != nil {
		return OrderLine{}, err
	}
	priceVo, err := NewOrderLinePrice(price)
	if err != nil {
		return OrderLine{}, err
	}
	quantityVo, err := NewOrderLineQuantity(quantity)
	if err != nil {
		return OrderLine{}, err
	}

	return OrderLine{
		ID:        idVo,
		ProductID: productIdVo,
		Price:     priceVo,
		Quantity:  quantityVo,
	}, nil
}
