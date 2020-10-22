package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

//InventorySyncPublisher struct
type InventorySyncPublisher struct {
	inventorySyncTopic *pubsub.Topic
	context            context.Context
}

//NewInventorySyncPublisher constructor
func NewInventorySyncPublisher(ctx context.Context, t *pubsub.Topic) Publisher {
	return &InventorySyncPublisher{inventorySyncTopic: t, context: ctx}
}

//Publish sends message to pubsub topic
func (p *InventorySyncPublisher) Publish(payload []byte, headers map[string]string) {
	result := p.inventorySyncTopic.Publish(p.context, &pubsub.Message{Data: payload, Attributes: headers})

	go func(res *pubsub.PublishResult, data []byte) {
		_, err := res.Get(p.context)
		if err != nil {
			fmt.Printf("Failed to publish: %v", err)
		} else {
			fmt.Printf("Message published: %s\n", string(data))
		}
	}(result, payload)
}
