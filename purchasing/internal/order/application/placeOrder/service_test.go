package creator

import (
	"context"
	"testing"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/application/mocks"
	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus/eventmocks"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_PlaceOrderService(t *testing.T) {
	ctx := context.Background()
	id := vo.NewID()
	addressVo, _ := domain.NewOrderAddressShipping("Indianápolis, Indiana 46218, EE. UU.")
	priceVo, _ := domain.NewOrderLinePrice(34.0)
	quantityVo, _ := domain.NewOrderLineQuantity(3)
	stateVo, _ := domain.NewOrderState(domain.ORDER_CREATED)
	order := domain.Order{
		ID:              vo.NewID(),
		ClientID:        vo.NewID(),
		AddressShipping: addressVo,
		State:           stateVo,
		OrderLines: []domain.OrderLine{
			{
				ID:        vo.NewID(),
				ProductID: vo.NewID(),
				Price:     priceVo,
				Quantity:  quantityVo,
			},
		},
	}

	orderRepositoryMock := mocks.NewOrderRepositoryMock()
	orderRepositoryMock.On("Search", mock.Anything, mock.AnythingOfType("vo.ID")).Return(order, nil).Once()
	orderRepositoryMock.On("BookProductFromInventory", mock.Anything, mock.AnythingOfType("vo.ID"), mock.AnythingOfType("domain.OrderLineQuantity")).Return(vo.NewID(), nil).Once()
	orderRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Order")).Return(nil).Once()

	eventBusMock := eventmocks.NewMockBus()
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]eventbus.Event")).Return(nil).Once().Run(func(args mock.Arguments) {
		events := args.Get(1).([]eventbus.Event)
		assert.Equal(t, 1, len(events))
	})

	placeOrderService := NewPlaceOrderService(orderRepositoryMock, eventBusMock)
	err := placeOrderService.Do(ctx, id)
	orderRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}
