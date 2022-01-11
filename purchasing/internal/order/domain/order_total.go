package domain

import "errors"

var ErrEmptyOrderTotal = errors.New("the field Order Total can not be 0")

type OrderTotal struct {
	value float64
}

func NewOrderTotal(value float64) (OrderTotal, error) {
	if value == 0 {
		return OrderTotal{}, ErrEmptyOrderTotal
	}

	return OrderTotal{
		value: value,
	}, nil
}

func (o OrderTotal) Primitive() float64 {
	return o.value
}
