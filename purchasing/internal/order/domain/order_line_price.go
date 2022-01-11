package domain

import "errors"

var ErrEmptyOrderLinePrice = errors.New("the field Order Line Price can not be 0")

type OrderLinePrice struct {
	value float64
}

func NewOrderLinePrice(value float64) (OrderLinePrice, error) {
	if value == 0 {
		return OrderLinePrice{}, ErrEmptyOrderLinePrice
	}

	return OrderLinePrice{
		value: value,
	}, nil
}

func (o OrderLinePrice) Primitive() float64 {
	return o.value
}
