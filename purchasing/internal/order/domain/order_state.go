package domain

const (
	ORDER_CREATED = "CREATED"
	ORDER_PLACED  = "PLACED"
)

type OrderState struct {
	value string
}

func NewOrderState() OrderState {
	return OrderState{
		value: ORDER_CREATED,
	}
}

func (o *OrderState) ChangeToPlaced() {
	o.value = ORDER_PLACED
}
