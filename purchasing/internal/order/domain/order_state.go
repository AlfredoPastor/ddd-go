package domain

import (
	"fmt"
)

const (
	ORDER_CREATED = "CREATED"
	ORDER_PLACED  = "PLACED"
)

type OrderState struct {
	value string
}

func NewOrderState(value string) (OrderState, error) {
	switch value {
	case ORDER_CREATED:
		return OrderState{
			value: ORDER_CREATED,
		}, nil
	case ORDER_PLACED:
		return OrderState{
			value: ORDER_PLACED,
		}, nil
	default:
		return OrderState{}, fmt.Errorf("state %s dosen't exist", value)
	}

}

func (o *OrderState) ChangeToPlaced() {
	o.value = ORDER_PLACED
}
