package domain

import "errors"

var ErrEmptyOrderAddressShipping = errors.New("the field Order Address Shipping can not be empty")

type OrderAddressShipping struct {
	value string
}

func NewOrderAddressShipping(address string) (OrderAddressShipping, error) {
	if address == "" {
		return OrderAddressShipping{}, ErrEmptyOrderAddressShipping
	}

	return OrderAddressShipping{
		value: address,
	}, nil
}

func (o OrderAddressShipping) Primitive() string {
	return o.value
}
