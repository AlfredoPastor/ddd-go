package domain

import "errors"

var ErrEmptyOrderLineQuantity = errors.New("the field Order Line Quantity can not be 0")

type OrderLineQuantity struct {
	value int
}

func NewOrderLineQuantity(value int) (OrderLineQuantity, error) {
	if value == 0 {
		return OrderLineQuantity{}, ErrEmptyOrderLineQuantity
	}

	return OrderLineQuantity{
		value: value,
	}, nil
}

func (o OrderLineQuantity) Primitive() int {
	return o.value
}
