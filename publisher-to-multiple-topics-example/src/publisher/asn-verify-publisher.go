package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

//AsnVerifyPublisher struct
type AsnVerifyPublisher struct {
	asnVerifyTopic *pubsub.Topic
	context        context.Context
}

//NewAsnVerifyPublisher constructor
func NewAsnVerifyPublisher(ctx context.Context, t *pubsub.Topic) Publisher {
	return &AsnVerifyPublisher{asnVerifyTopic: t, context: ctx}
}

//Publish sends message to pubsub topic
func (p *AsnVerifyPublisher) Publish(payload []byte, headers map[string]string) {
	result := p.asnVerifyTopic.Publish(p.context, &pubsub.Message{Data: payload, Attributes: headers})

	go func(res *pubsub.PublishResult, data []byte) {
		_, err := res.Get(p.context)
		if err != nil {
			fmt.Printf("Failed to publish: %v", err)
		} else {
			fmt.Printf("Message published: %s\n", string(data))
		}
	}(result, payload)
}
