package store

import (
	"log"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderAdapter struct {
	ID              primitive.ObjectID `bson:"_id"`
	ClientID        primitive.ObjectID `bson:"client_id"`
	AddressShipping string             `bson:"address_shipping"`
	State           string             `bson:"state"`
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
	orderAdapter.CompleteOrderLines(order.OrderLines)
	return orderAdapter
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

func NewOrderFromDatabase(orderAdapter OrderAdapter) (domain.Order, error) {
	idVo, err := vo.NewIDFromString(orderAdapter.ID.Hex())
	if err != nil {
		return domain.Order{}, err
	}
	idClientVo, err := vo.NewIDFromString(orderAdapter.ClientID.Hex())
	if err != nil {
		return domain.Order{}, err
	}
	addressVo, err := domain.NewOrderAddressShipping(orderAdapter.AddressShipping)
	if err != nil {
		return domain.Order{}, err
	}
	stateVo, err := domain.NewOrderState(orderAdapter.State)
	if err != nil {
		return domain.Order{}, err
	}
	taxesVo, err := domain.NewOrderTaxes(orderAdapter.Taxes)
	if err != nil {
		return domain.Order{}, err
	}
	subtotalVo, err := domain.NewOrderSubtotal(orderAdapter.Subtotal)
	if err != nil {
		return domain.Order{}, err
	}
	totalVo, err := domain.NewOrderTotal(orderAdapter.Total)
	if err != nil {
		return domain.Order{}, err
	}
	order := domain.Order{
		ID:              idVo,
		ClientID:        idClientVo,
		AddressShipping: addressVo,
		State:           stateVo,
		Taxes:           taxesVo,
		Subtotal:        subtotalVo,
		Total:           totalVo,
		OrderLines:      []domain.OrderLine{},
	}
	for _, orderLineAdapter := range orderAdapter.OrderLines {
		idVo, err := vo.NewIDFromString(orderLineAdapter.ID.Hex())
		if err != nil {
			return domain.Order{}, err
		}
		idProductVo, err := vo.NewIDFromString(orderLineAdapter.ProductID.Hex())
		if err != nil {
			return domain.Order{}, err
		}
		priceVo, err := domain.NewOrderLinePrice(orderLineAdapter.Price)
		if err != nil {
			return domain.Order{}, err
		}
		quantityVo, err := domain.NewOrderLineQuantity(orderLineAdapter.Quantity)
		if err != nil {
			return domain.Order{}, err
		}
		orderLine := domain.OrderLine{
			ID:        idVo,
			ProductID: idProductVo,
			Price:     priceVo,
			Quantity:  quantityVo,
		}
		order.OrderLines = append(order.OrderLines, orderLine)
	}

	return order, nil
}
