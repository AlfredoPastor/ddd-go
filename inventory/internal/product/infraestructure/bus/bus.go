package bus

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/AlfredoPastor/ddd-go/shared/eventbus"
	"github.com/AlfredoPastor/ddd-go/shared/rabbitmq"
)

type Bus struct {
	*rabbitmq.RabbitClient
}

func (b Bus) Publish(ctx context.Context, events []eventbus.Event) error {
	for _, ev := range events {
		if ctx.Err() == context.Canceled {
			return errors.New("petición cancelada")
		}
		message, err := json.Marshal(ev)
		if err != nil {
			return err
		}
		b.RabbitClient.Publisher(string(ev.Type()), message)
	}
	return nil
}

func (b Bus) Subscribe(ctx context.Context, topic eventbus.Type, evt eventbus.Handler) {
	go b.RabbitClient.Subscriber(ctx, string(topic), evt.Handle)
}

func (b Bus) Run(ctx context.Context) error {
	if err := b.RabbitClient.Run(ctx); err != nil {
		return err
	}
	return nil
}

func NewBus(cli *rabbitmq.RabbitClient) Bus {
	return Bus{
		RabbitClient: cli,
	}
}
