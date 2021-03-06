package domain

import (
	"encoding/json"
	"fmt"
)

type OrderAdapter struct {
	ID              string             `json:"id"`
	ClientID        string             `json:"client_id"`
	AddressShipping string             `json:"address_shipping"`
	Taxes           float64            `json:"taxes"`
	Subtotal        float64            `json:"subtotal"`
	Total           float64            `json:"total"`
	OrderLines      []OrderLineAdapter `json:"order_lines"`
}

type OrderLineAdapter struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

func NewOrderAdapter(order Order) OrderAdapter {
	orderAdapter := OrderAdapter{
		ID:              order.ID.String(),
		ClientID:        order.ClientID.String(),
		AddressShipping: order.AddressShipping.Primitive(),
		Taxes:           order.Taxes.Primitive(),
		Subtotal:        order.Subtotal.Primitive(),
		Total:           order.Total.Primitive(),
	}
	orderAdapter.CompleteOrderLines(order.OrderLines)
	return orderAdapter
}

func (o *OrderAdapter) Serialize() []byte {
	entity, err := json.Marshal(o)
	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return []byte{}
	}

	return entity
}

func (o *OrderAdapter) CompleteOrderLines(orderLines []OrderLine) {
	o.OrderLines = []OrderLineAdapter{}
	for _, orderLine := range orderLines {
		orderLineAdapter := OrderLineAdapter{
			ID:        orderLine.ID.String(),
			ProductID: orderLine.ID.String(),
			Price:     orderLine.Price.Primitive(),
			Quantity:  orderLine.Quantity.Primitive(),
		}
		o.OrderLines = append(o.OrderLines, orderLineAdapter)
	}
}
