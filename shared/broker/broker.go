package domain

import (
	"context"
	"sync"
)

type Broker struct {
	mu     sync.Mutex
	topics map[string][]Subscriber
}

func (b *Broker) AddSubscriber(topic string, ch Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.topics[topic]; ok {
		b.topics[topic] = append(b.topics[topic], ch)
	} else {
		b.topics[topic] = []Subscriber{ch}
	}
}

func (b *Broker) Publish(ctx context.Context, topic string, message []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.topics[topic]; ok {
		for _, handler := range b.topics[topic] {
			go func(handler Subscriber) {
				handler <- message
			}(handler)
		}
	}
}

func (b *Broker) Subscribe(ctx context.Context, topic string, function func(interface{})) {
	sub := NewSubscriber()
	b.AddSubscriber(topic, sub)
	for {
		function(<-sub)
	}
}

func NewBroker() *Broker {
	return &Broker{
		topics: make(map[string][]Subscriber),
	}
}
