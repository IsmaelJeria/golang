package listener

import (
	"context"
	"internal-listener/listener/gscdictionary"
	"internal-listener/rest"
	"log"

	"cloud.google.com/go/pubsub"
	guuid "github.com/google/uuid"
)

//PriceInboundTopicGSCListener struct
type PriceInboundTopicGSCListener struct {
	sub         *pubsub.Subscription
	PriceBFF    *rest.PriceBFF
	priceLogACL *rest.PriceLogACL
}

//NewPriceInboundTopicGSCListener PriceInboundTopicGSCListener struct constructor
func NewPriceInboundTopicGSCListener(s *pubsub.Subscription, p *rest.PriceBFF, pl *rest.PriceLogACL) *PriceInboundTopicGSCListener {
	return &PriceInboundTopicGSCListener{sub: s, PriceBFF: p, priceLogACL: pl}
}

//PullMsgs method pull price-inbound-topic-GSC messages
func (l *PriceInboundTopicGSCListener) PullMsgs() {
	err := l.sub.Receive(context.Background(), func(ctx context.Context, message *pubsub.Message) {
		message.Ack()
		log.Printf("message received from GSC: %s\n", message.Data)
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
		m["x-channel"] = "GSC"
		m["x-txref"] = guuid.New().String()
		var productPrice *gscdictionary.ProductPrice
		if message.Attributes["eventtype"] == "priceCreated" {
			destinyData := productPrice.GscPOSTToProductPrice(message.Data)
			l.PriceBFF.Update(destinyData, m)
		}
		if message.Attributes["eventtype"] == "createProductPrice" {
			destinyData := productPrice.GscPATCHToProductPrice(message.Data)
			l.priceLogACL.Update(destinyData, m)
		}
		if message.Attributes["eventtype"] == "createProductPriceError" {
			l.priceLogACL.Delete(message.Data, m)
		}

	})

	if err != nil {
		log.Printf("Error :  %s", err)
	}
}
