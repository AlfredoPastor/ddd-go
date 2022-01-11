package creator

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func NewOrderRepositoryMock() *OrderRepositoryMock {
	return &OrderRepositoryMock{}
}

func (o *OrderRepositoryMock) SearchByClient(ctx context.Context, id vo.ID) (domain.Order, error) {
	args := o.Called(ctx, id)
	return args.Get(0).(domain.Order), args.Error(1)
}

func (o *OrderRepositoryMock) Search(ctx context.Context, id vo.ID) (domain.Order, error) {
	args := o.Called(ctx, id)
	return args.Get(0).(domain.Order), args.Error(1)
}

func (o *OrderRepositoryMock) Delete(ctx context.Context, id vo.ID) error {
	args := o.Called(ctx, id)
	return args.Error(0)
}

func (o *OrderRepositoryMock) Save(ctx context.Context, order domain.Order) error {
	args := o.Called(ctx, order)
	return args.Error(0)
}
