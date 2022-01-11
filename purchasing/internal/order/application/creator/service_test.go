package creator

import (
	"context"
	"errors"
	"testing"

	"github.com/AlfredoPastor/ddd-go/purchasing/internal/order/domain"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/eventbus/eventmocks"
	"github.com/AlfredoPastor/ddd-go/shared/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type OrderTest struct {
	ID              string
	ClientID        string
	AddressShipping string
	OrderLines      []domain.OrderLine
}

func NewOrderTest() OrderTest {
	id := uuid.NewID()
	clientId := uuid.NewID()
	addressShipping := "Indianápolis, Indiana 46218, EE. UU."

	orderTest := OrderTest{
		ID:              id,
		ClientID:        clientId,
		AddressShipping: addressShipping,
	}

	orderTest.OrderLines = []domain.OrderLine{}

	idOrderLine := uuid.NewID()
	productIdOrderLine := uuid.NewID()
	priceOrderLine := 34.0
	quantityOrderLine := 1
	orderLine, _ := domain.NewOrderLine(idOrderLine, productIdOrderLine, priceOrderLine, quantityOrderLine)
	orderTest.OrderLines = append(orderTest.OrderLines, orderLine)
	return orderTest
}

func Test_OrderCreatorService(t *testing.T) {
	ctx := context.Background()
	orderTest := NewOrderTest()

	orderRepositoryMock := NewOrderRepositoryMock()
	orderRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Order")).Return(nil).Once()

	eventBusMock := eventmocks.NewMockBus()
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]eventbus.Event")).Return(nil).Once().Run(func(args mock.Arguments) {
		events := args.Get(1).([]eventbus.Event)
		assert.Equal(t, len(events), 1)
	})

	orderCreatorService := NewOrderCreatorService(orderRepositoryMock, eventBusMock)
	err := orderCreatorService.Create(ctx, orderTest.ID, orderTest.ClientID, orderTest.AddressShipping, orderTest.OrderLines)
	orderRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_OrderCreatorService_Save_Error(t *testing.T) {
	ctx := context.Background()
	orderTest := NewOrderTest()

	eventBusMock := eventmocks.NewMockBus()
	orderRepositoryMock := NewOrderRepositoryMock()
	orderRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Order")).Return(errors.New("something unexpected happened")).Once()

	orderCreatorService := NewOrderCreatorService(orderRepositoryMock, eventBusMock)
	err := orderCreatorService.Create(ctx, orderTest.ID, orderTest.ClientID, orderTest.AddressShipping, orderTest.OrderLines)
	orderRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_OrderCreatorService_OrderEntityError(t *testing.T) {
	ctx := context.Background()
	orderTest := NewOrderTest()
	orderTest.AddressShipping = ""

	eventBusMock := eventmocks.NewMockBus()
	orderRepositoryMock := NewOrderRepositoryMock()

	orderCreatorService := NewOrderCreatorService(orderRepositoryMock, eventBusMock)
	err := orderCreatorService.Create(ctx, orderTest.ID, orderTest.ClientID, orderTest.AddressShipping, orderTest.OrderLines)
	orderRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}

func Test_OrderCreatorService_PublishEventError(t *testing.T) {
	ctx := context.Background()
	orderTest := NewOrderTest()

	orderRepositoryMock := NewOrderRepositoryMock()
	orderRepositoryMock.On("Save", mock.Anything, mock.AnythingOfType("domain.Order")).Return(nil).Once()

	eventBusMock := eventmocks.NewMockBus()
	eventBusMock.On("Publish", mock.Anything, mock.AnythingOfType("[]eventbus.Event")).Return(errors.New("something unexpected happened")).Once().Run(func(args mock.Arguments) {
		events := args.Get(1).([]eventbus.Event)
		assert.Equal(t, len(events), 1)
	})

	orderCreatorService := NewOrderCreatorService(orderRepositoryMock, eventBusMock)
	err := orderCreatorService.Create(ctx, orderTest.ID, orderTest.ClientID, orderTest.AddressShipping, orderTest.OrderLines)
	orderRepositoryMock.AssertExpectations(t)
	eventBusMock.AssertExpectations(t)
	assert.Error(t, err)
}
