package domain

import "errors"

var ErrEmptyOrderSubtotal = errors.New("the field Order Subtotal can not be 0")

type OrderSubtotal struct {
	value float64
}

func NewOrderSubtotal(value float64) (OrderSubtotal, error) {
	if value == 0 {
		return OrderSubtotal{}, ErrEmptyOrderSubtotal
	}

	return OrderSubtotal{
		value: value,
	}, nil
}

func (o OrderSubtotal) Primitive() float64 {
	return o.value
}
