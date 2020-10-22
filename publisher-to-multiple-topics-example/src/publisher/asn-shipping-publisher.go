package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

//ASNShippingPublisher struct
type ASNShippingPublisher struct {
	asnShippingTopic *pubsub.Topic
	context          context.Context
}

//NewASNShippingPublisher constructor
func NewASNShippingPublisher(ctx context.Context, t *pubsub.Topic) Publisher {
	return &ASNShippingPublisher{asnShippingTopic: t, context: ctx}
}

//Publish sends message to pubsub topic
func (p *ASNShippingPublisher) Publish(payload []byte, headers map[string]string) {
	result := p.asnShippingTopic.Publish(p.context, &pubsub.Message{Data: payload, Attributes: headers})

	go func(res *pubsub.PublishResult, data []byte) {
		_, err := res.Get(p.context)
		if err != nil {
			fmt.Printf("Failed to publish: %v", err)
		} else {
			fmt.Printf("Message published: %s\n", string(data))
		}
	}(result, payload)
}
