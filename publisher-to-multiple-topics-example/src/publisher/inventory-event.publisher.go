package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

//InventoryEventPublisher struct
type InventoryEventPublisher struct {
	inventoryEventTopic *pubsub.Topic
	context             context.Context
}

//NewInventoryEventPublisher constructor
func NewInventoryEventPublisher(ctx context.Context, t *pubsub.Topic) Publisher {
	return &InventoryEventPublisher{inventoryEventTopic: t, context: ctx}
}

//Publish sends message to pubsub topic
func (p *InventoryEventPublisher) Publish(payload []byte, headers map[string]string) {
	result := p.inventoryEventTopic.Publish(p.context, &pubsub.Message{Data: payload, Attributes: headers})

	go func(res *pubsub.PublishResult, data []byte) {
		_, err := res.Get(p.context)
		if err != nil {
			fmt.Printf("Failed to publish: %v", err)
		} else {
			fmt.Printf("Message published: %s\n", string(data))
		}
	}(result, payload)
}
