package listener

import (
	"context"
	"fmt"
	"internal-listener/rest"
	"log"

	"cloud.google.com/go/pubsub"
)

//PriceInternalTopicWebListener struct
type PriceInternalTopicWebListener struct {
	sub         *pubsub.Subscription
	priceLogACL *rest.PriceLogACL
}

//NewPriceInternalTopicWebListener PriceInternalTopicWebListener struct constructor
func NewPriceInternalTopicWebListener(s *pubsub.Subscription, p *rest.PriceLogACL) *PriceInternalTopicWebListener {
	return &PriceInternalTopicWebListener{sub: s, priceLogACL: p}
}

//PullMsgs method pull price-internal-topic-web messages
func (l *PriceInternalTopicWebListener) PullMsgs() {
	err := l.sub.Receive(context.Background(), func(ctx context.Context, message *pubsub.Message) {
		message.Ack()
		fmt.Println(message.Attributes)
		m := make(map[string]string)
		m["x-usrtx"] = message.Attributes["x-usrtx"]
		m["accept"] = "application/json"
		m["x-rhsref"] = "default"
		m["x-prref"] = "default"
		m["x-isindividual"] = message.Attributes["isIndividual"]
		m["x-filename"] = ""
		m["x-country"] = "CL"
		m["x-commerce"] = "Falabella"
		m["x-cmref"] = "Ecommerce"
		m["x-chref"] = "WEB"
		m["x-channel"] = "WEB"
		m["x-txref"] = message.Attributes["x-txref"]
		l.priceLogACL.Save(message.Data, m)
		log.Printf("message received from web: %s\n", message.Data)
	})

	if err != nil {
		log.Printf("Error :  %s", err)
	}
}
