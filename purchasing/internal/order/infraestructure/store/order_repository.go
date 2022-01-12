package store

import (
	"context"
	"time"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/shared/mongoatlas"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	*mongoatlas.MongoClient
	collection *mongo.Collection
}

func NewOrderRepository(client *mongoatlas.MongoClient) OrderRepository {
	return OrderRepository{
		MongoClient: client,
		collection:  client.CreateCollection("order"),
	}
}

func (o OrderRepository) BookProductFromInventory(ctx context.Context, id vo.ID, quantity domain.OrderLineQuantity) (vo.ID, error) {
	return vo.ID{}, nil
}

func (o OrderRepository) SearchByClient(ctx context.Context, clientID vo.ID) (domain.Order, error) {
	return domain.Order{}, nil
}

func (o OrderRepository) Search(ctx context.Context, id vo.ID) (domain.Order, error) {
	return domain.Order{}, nil
}

func (o OrderRepository) Delete(ctx context.Context, id vo.ID) error {
	return nil
}

func (o OrderRepository) Save(ctx context.Context, order domain.Order) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	orderAdapter := NewOrderAdapter(order)
	_, err := o.collection.InsertOne(ctx, orderAdapter)
	if err != nil {
		return err
	}
	return nil
}
