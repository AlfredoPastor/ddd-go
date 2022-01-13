package substractQuantity

import (
	"context"
	"testing"

	"github.com/AlfredoPastor/ddd-go/inventory/internal/product/application/mocks"
	"github.com/AlfredoPastor/ddd-go/inventory/internal/product/domain"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus/eventmocks"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_SubstractQuantityService(t *testing.T) {
	ctx := context.Background()
	productID := vo.NewID()
	quantity := domain.NewProductQuantity(2)

	product := domain.Product{
		ID:       productID,
		Name:     domain.NewProductName("Ball"),
		Quantity: domain.NewProductQuantity(4),
	}

	productRepositoryMock := mocks.NewProductRepositoryMock()
	productRepositoryMock.On("Search", mock.Anything, mock.AnythingOfType("vo.ID")).Return(product, nil).Once()
	productRepositoryMock.On("Update", mock.Anything, mock.AnythingOfType("domain.Product")).Return(nil).Once().Run(func(args mock.Arguments) {
		product := args.Get(1).(domain.Product)
		assert.Equal(t, 2, product.Quantity.Primitive())
	})

	eventBusMock := eventmocks.NewMockBus()
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]eventbus.Event")).Return(nil).Once().Run(func(args mock.Arguments) {
		events := args.Get(1).([]eventbus.Event)
		assert.Equal(t, 0, len(events))
	})
	substractQuantityService := NewSubstractQuantityService(productRepositoryMock, eventBusMock)
	err := substractQuantityService.Subtract(ctx, productID, quantity)

	productRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
