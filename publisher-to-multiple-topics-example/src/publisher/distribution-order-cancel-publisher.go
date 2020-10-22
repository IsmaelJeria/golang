package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

//DOCancelPublisher struct
type DOCancelPublisher struct {
	doCancelTopic *pubsub.Topic
	context       context.Context
}

//NewDOCancelPublisher constructor
func NewDOCancelPublisher(ctx context.Context, t *pubsub.Topic) Publisher {
	return &DOCancelPublisher{doCancelTopic: t, context: ctx}
}

//Publish sends message to pubsub topic
func (p *DOCancelPublisher) Publish(payload []byte, headers map[string]string) {
	result := p.doCancelTopic.Publish(p.context, &pubsub.Message{Data: payload, Attributes: headers})

	go func(res *pubsub.PublishResult, data []byte) {
		_, err := res.Get(p.context)
		if err != nil {
			fmt.Printf("Failed to publish: %v", err)
		} else {
			fmt.Printf("Message published: %s\n", string(data))
		}
	}(result, payload)
}
