package eventmocks

import (
	"context"

	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/stretchr/testify/mock"
)

type MockBus struct {
	mock.Mock
}

func NewMockBus() *MockBus {
	return &MockBus{}
}

func (b *MockBus) Publish(ctx context.Context, events []eventbus.Event) error {
	args := b.Called(ctx, events)
	return args.Error(0)
}

func (b *MockBus) Subscribe(ctx context.Context, tipo eventbus.Type, handler eventbus.Handler) {
	b.Called(tipo, handler)
}
