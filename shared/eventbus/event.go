package eventbus

import (
	"context"
	"time"

	"github.com/AlfredoPastor/ddd-go/shared/uuid"
	"github.com/AlfredoPastor/ddd-go/shared/vo"
)

type Handler interface {
	Handle([]byte) bool
}

// Bus defines the expected behaviour from an event bus.
type Bus interface {
	// Publish is the method used to publish new events.
	Publish(context.Context, []Event) error
	// Subscribe is the method used to subscribe new event handlers.
	Subscribe(context.Context, Type, Handler)
}

// Type represents a domain event type.
type Body []byte

// Type represents a domain event type.
type Type string

// Event represents a domain command.
type Event interface {
	Type() Type
}

type BaseEvent struct {
	EventID     string
	AggregateID string
	OccurredOn  time.Time
}

func NewBaseEvent(aggregateID string) BaseEvent {
	return BaseEvent{
		EventID:     uuid.NewGoogleUuid(),
		AggregateID: aggregateID,
		OccurredOn:  time.Now(),
	}
}

func (e *BaseEvent) IsNotValid() bool {
	_, err := vo.NewIDFromString(e.AggregateID)
	if err != nil {
		return true
	}
	_, err = uuid.GoogleUuidParse(e.EventID)
	if err != nil {
		return true
	}
	return err != nil
}
