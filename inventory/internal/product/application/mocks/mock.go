package mocks

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/inventory/internal/product/domain"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func NewProductRepositoryMock() ProductRepositoryMock {
	return ProductRepositoryMock{}
}

func (p ProductRepositoryMock) Search(ctx context.Context, id vo.ID) (domain.Product, error) {
	args := p.Called(ctx, id)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (p ProductRepositoryMock) Update(ctx context.Context, product domain.Product) error {
	args := p.Called(ctx, product)
	return args.Error(0)
}
