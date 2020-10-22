package publisher

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
)

//DistributionOrderUpdatePublisher struct
type DistributionOrderUpdatePublisher struct {
	distributionOrderUpdateTopic *pubsub.Topic
	context                      context.Context
}

//NewDistributionOrderUpdatePublisher constructor
func NewDistributionOrderUpdatePublisher(ctx context.Context, t *pubsub.Topic) Publisher {
	return &DistributionOrderUpdatePublisher{distributionOrderUpdateTopic: t, context: ctx}
}

//Publish sends message to pubsub topic
func (p *DistributionOrderUpdatePublisher) Publish(payload []byte, headers map[string]string) {
	result := p.distributionOrderUpdateTopic.Publish(p.context, &pubsub.Message{Data: payload, Attributes: headers})

	go func(res *pubsub.PublishResult, data []byte) {
		_, err := res.Get(p.context)
		if err != nil {
			fmt.Printf("Failed to publish: %v", err)
		} else {
			fmt.Printf("Message published: %s\n", string(data))
		}
	}(result, payload)
}
