package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

//ExamplePublisher struct
type ExamplePublisher struct {
	examplePublisher *pubsub.Topic
	context          context.Context
}

//NewExamplePublisher constructor
func NewExamplePublisher(ctx context.Context, t *pubsub.Topic) Publisher {
	return &ExamplePublisher{examplePublisher: t, context: ctx}
}

//Publish sends message to pubsub topic
func (p *ExamplePublisher) Publish(payload []byte, headers map[string]string) {
	result := p.examplePublisher.Publish(p.context, &pubsub.Message{Data: payload, Attributes: headers})
	go func(res *pubsub.PublishResult, data []byte) {
		_, err := res.Get(p.context)
		if err != nil {
			fmt.Printf("Failed to publish: %v", err)
		} else {
			fmt.Printf("Message published: %s\n", string(data))
		}
	}(result, payload)
}
