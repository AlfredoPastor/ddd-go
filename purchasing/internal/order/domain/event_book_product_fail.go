package domain

import (
	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
)

const ProductBookedFailEventType eventbus.Type = "purchasing.1.event.product.booked.fail"

type ProductBookedFailEvent struct {
	eventbus.BaseEvent
}

func NewProductBookedFailEvent(idBooked vo.ID) ProductBookedFailEvent {
	return ProductBookedFailEvent{
		BaseEvent: eventbus.NewBaseEvent(idBooked.String()),
	}
}

func (e ProductBookedFailEvent) Type() eventbus.Type {
	return ProductBookedFailEventType
}
