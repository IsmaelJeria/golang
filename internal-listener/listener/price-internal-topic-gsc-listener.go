package listener

import (
	"context"
	"fmt"
	"internal-listener/rest"
	"log"

	"cloud.google.com/go/pubsub"
	guuid "github.com/google/uuid"
)

//PriceInternalTopicGSCListener struct
type PriceInternalTopicGSCListener struct {
	sub         *pubsub.Subscription
	priceLogACL *rest.PriceLogACL
}

//NewPriceInternalTopicGSCListener PriceInternalTopicGSCListener struct constructor
func NewPriceInternalTopicGSCListener(s *pubsub.Subscription, p *rest.PriceLogACL) *PriceInternalTopicGSCListener {
	return &PriceInternalTopicGSCListener{sub: s, priceLogACL: p}
}

//PullMsgs method pull price-internal-topic-GSC messages
func (l *PriceInternalTopicGSCListener) PullMsgs() {
	err := l.sub.Receive(context.Background(), func(ctx context.Context, message *pubsub.Message) {
		message.Ack()
		fmt.Printf("headers received from GSC: %s\n", message.Attributes)
		m := make(map[string]string)
		m["x-usrtx"] = "CATALYST"
		m["accept"] = "application/json"
		m["x-rhsref"] = "default"
		m["x-prref"] = "default"
		m["x-isindividual"] = "true"
		m["x-filename"] = "NOFILE"
		m["x-country"] = "CL"
		m["x-commerce"] = "Falabella"
		m["x-cmref"] = "Ecommerce"
		m["x-chref"] = "GSC"
		m["x-txref"] = guuid.New().String()
		l.priceLogACL.Save(message.Data, m)
		log.Printf("message received from GSC: %s\n", message.Data)
	})

	if err != nil {
		log.Printf("Error :  %s", err)
	}
}
