package store

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/shared/mongoatlas"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	*mongoatlas.MongoClient
	collection *mongo.Collection
}

func NewProductRepository(client *mongoatlas.MongoClient) ProductRepository {
	return ProductRepository{
		MongoClient: client,
		collection:  client.CreateCollection("product"),
	}
}

func (o ProductRepository) BookProductFromInventory(ctx context.Context) (vo.ID, error) {
	return vo.ID{}, nil
}

func (o ProductRepository) SearchByClient(ctx context.Context) error {
	return nil
}

func (o ProductRepository) Search(ctx context.Context, id vo.ID) error {
	return nil
}

func (o ProductRepository) Delete(ctx context.Context, id vo.ID) error {
	return nil
}

func (o ProductRepository) Save(ctx context.Context) error {

	return nil
}
