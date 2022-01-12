package store

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderAdapter struct {
	ID              primitive.ObjectID `bson:"_id"`
	ClientID        primitive.ObjectID `bson:"client_id"`
	AddressShipping string             `bson:"address_shipping"`
	Taxes           float64            `bson:"taxes"`
	Subtotal        float64            `bson:"subtotal"`
	Total           float64            `bson:"total"`
	OrderLines      []OrderLineAdapter `bson:"order_lines"`
}

type OrderLineAdapter struct {
	ID        primitive.ObjectID `bson:"id"`
	ProductID primitive.ObjectID `bson:"product_id"`
	Price     float64            `bson:"price"`
	Quantity  int                `bson:"quantity"`
}

func NewOrderAdapter(order domain.Order) OrderAdapter {
	id, err := primitive.ObjectIDFromHex(order.ID.String())
	if err != nil {
		log.Println(err.Error())
	}
	clientId, err := primitive.ObjectIDFromHex(order.ClientID.String())
	if err != nil {
		log.Println(err.Error())
	}
	orderAdapter := OrderAdapter{
		ID:              id,
		ClientID:        clientId,
		AddressShipping: order.AddressShipping.Primitive(),
		Taxes:           order.Taxes.Primitive(),
		Subtotal:        order.Subtotal.Primitive(),
		Total:           order.Total.Primitive(),
	}
	orderAdapter.CompleteOrderLines(order.OderLines)
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

func (o *OrderAdapter) CompleteOrderLines(orderLines []domain.OrderLine) {
	o.OrderLines = []OrderLineAdapter{}
	for _, orderLine := range orderLines {
		id, err := primitive.ObjectIDFromHex(orderLine.ID.String())
		if err != nil {
			log.Println(err.Error())
		}
		productId, err := primitive.ObjectIDFromHex(orderLine.ProductID.String())
		if err != nil {
			log.Println(err.Error())
		}
		orderLineAdapter := OrderLineAdapter{
			ID:        id,
			ProductID: productId,
			Price:     orderLine.Price.Primitive(),
			Quantity:  orderLine.Quantity.Primitive(),
		}
		o.OrderLines = append(o.OrderLines, orderLineAdapter)
	}
}

func NewOrderAdapterDeserialize(data []byte) (domain.Order, error) {
	orderAdapter := OrderAdapter{}
	err := json.Unmarshal(data, &orderAdapter)
	if err != nil {
		return domain.Order{}, err
	}
	orderLines := []domain.OrderLine{}
	for _, orderLineAdapter := range orderAdapter.OrderLines {
		orderLine, err := domain.NewOrderLine(orderLineAdapter.ID.Hex(), orderLineAdapter.ProductID.Hex(), orderLineAdapter.Price, orderLineAdapter.Quantity)
		if err != nil {
			return domain.Order{}, err
		}
		orderLines = append(orderLines, orderLine)
	}
	order, err := domain.NewOrder(orderAdapter.ID.Hex(), orderAdapter.ClientID.Hex(), orderAdapter.AddressShipping, orderLines)
	if err != nil {
		return domain.Order{}, err
	}

	return order, nil
}
