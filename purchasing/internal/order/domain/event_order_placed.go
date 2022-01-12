package domain

import (
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
)

const OrderPlacedEventType eventbus.Type = "purchasing.1.event.order.placed"

type OrderPlacedEvent struct {
	eventbus.BaseEvent
	Entity eventbus.Body
}

func NewOrderPlacedEvent(order Order) OrderPlacedEvent {
	orderAdapter := NewOrderAdapter(order)
	entity := orderAdapter.Serialize()

	return OrderPlacedEvent{
		BaseEvent: eventbus.NewBaseEvent(order.ID.String()),
		Entity:    entity,
	}
}

func (e OrderPlacedEvent) Type() eventbus.Type {
	return OrderPlacedEventType
}

func (e OrderPlacedEvent) Body() eventbus.Body {
	return e.Entity
}
