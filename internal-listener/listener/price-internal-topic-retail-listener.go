package listener

import (
	"context"
	"internal-listener/rest"
	"log"

	"cloud.google.com/go/pubsub"
)

//PriceInternalTopicRetailListener struct
type PriceInternalTopicRetailListener struct {
	sub         *pubsub.Subscription
	priceLogACL *rest.PriceLogACL
}

//NewPriceInternalTopicRetailListener PriceInternalTopicRetailListener struct constructor
func NewPriceInternalTopicRetailListener(s *pubsub.Subscription, p *rest.PriceLogACL) *PriceInternalTopicRetailListener {
	return &PriceInternalTopicRetailListener{sub: s, priceLogACL: p}
}

//PullMsgs method pull price-internal-topic-retail messages
func (l *PriceInternalTopicRetailListener) PullMsgs() {
	err := l.sub.Receive(context.Background(), func(ctx context.Context, message *pubsub.Message) {
		message.Ack()
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
		m["x-chref"] = "RETAIL"
		m["x-channel"] = "RETAIL"
		m["x-txref"] = message.Attributes["x-txref"]
		l.priceLogACL.Save(message.Data, m)
		log.Printf("message received from retail: %s\n", message.Data)
	})

	if err != nil {
		log.Printf("Error :  %s", err)
	}
}
