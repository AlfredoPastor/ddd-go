package domain

import "errors"

var ErrEmptyOrderTaxes = errors.New("the field Order Taxes can not be 0")

type OrderTaxes struct {
	value float64
}

func NewOrderTaxes(value float64) (OrderTaxes, error) {
	if value == 0 {
		return OrderTaxes{}, ErrEmptyOrderTaxes
	}

	return OrderTaxes{
		value: value,
	}, nil
}

func (o OrderTaxes) Primitive() float64 {
	return o.value
}
