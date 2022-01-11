package domain

import (
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
)

const OrderCreatedEventType eventbus.Type = "purchasing.1.event.order.created"

type OrderCreatedEvent struct {
	eventbus.BaseEvent
	Entity eventbus.Body
}

func NewOrderCreatedEvent(order Order) OrderCreatedEvent {
	orderAdapter := NewOrderAdapter(order)
	entity := orderAdapter.Serialize()

	return OrderCreatedEvent{
		BaseEvent: eventbus.NewBaseEvent(order.ID.String()),
		Entity:    entity,
	}
}

func (e OrderCreatedEvent) Type() eventbus.Type {
	return OrderCreatedEventType
}

func (e OrderCreatedEvent) Body() eventbus.Body {
	return e.Entity
}
