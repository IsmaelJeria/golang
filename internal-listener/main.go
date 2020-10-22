package main

import (
	"context"
	"fmt"
	"internal-listener/listener"
	"internal-listener/rest"
	"log"
	"os"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

func loadEnviromentVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnviromentVariables()

	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	ctx := context.Background()

	pubSubClient, clientErr := pubsub.NewClient(ctx, projectID)
	if clientErr != nil {
		fmt.Println("Client error")
		// Handle error.
	}

	var wg sync.WaitGroup
	wg.Add(1)

	fastHTTPClient := &fasthttp.Client{}
	aclApigeeDNS := os.Getenv("PRICE_LOG_ACL_DNS")
	priceLogACL := rest.NewPriceLogACL(&aclApigeeDNS, fastHTTPClient)

	priceInternalTopicRetailSub := pubSubClient.Subscription(os.Getenv("PRICE_INTERNAL_TOPIC_RETAIL_SUBSCRIPTION"))
	priceInternalTopicRetailListener := listener.NewPriceInternalTopicRetailListener(priceInternalTopicRetailSub, priceLogACL)
	go priceInternalTopicRetailListener.PullMsgs()
	fmt.Println("pulling msgs from: " + os.Getenv("PRICE_INTERNAL_TOPIC_RETAIL_SUBSCRIPTION"))

	priceInternalTopicWebSub := pubSubClient.Subscription(os.Getenv("PRICE_INTERNAL_TOPIC_WEB_SUBSCRIPTION"))
	priceInternalTopicWebListener := listener.NewPriceInternalTopicWebListener(priceInternalTopicWebSub, priceLogACL)
	go priceInternalTopicWebListener.PullMsgs()
	fmt.Println("pulling msgs from: " + os.Getenv("PRICE_INTERNAL_TOPIC_WEB_SUBSCRIPTION"))

	priceInternalTopicGSCSub := pubSubClient.Subscription(os.Getenv("PRICE_INTERNAL_TOPIC_GSC_SUBSCRIPTION"))
	priceInternalTopicGSCListener := listener.NewPriceInternalTopicGSCListener(priceInternalTopicGSCSub, priceLogACL)
	go priceInternalTopicGSCListener.PullMsgs()
	fmt.Println("pulling msgs from: " + os.Getenv("PRICE_INTERNAL_TOPIC_GSC_SUBSCRIPTION"))

	bffDNS := os.Getenv("PRICE_BFF_DNS")
	priceBFF := rest.NewPriceBFF(&bffDNS, fastHTTPClient)
	priceInternalInboundTopicGSCSub := pubSubClient.Subscription(os.Getenv("PRICE_INBOUND_TOPIC_GSC_SUBSCRIPTION"))
	priceInternalInboundTopicGSCListener := listener.NewPriceInboundTopicGSCListener(priceInternalInboundTopicGSCSub, priceBFF, priceLogACL)
	go priceInternalInboundTopicGSCListener.PullMsgs()
	fmt.Println("pulling msgs from: " + os.Getenv("PRICE_INBOUND_TOPIC_GSC_SUBSCRIPTION"))

	wg.Wait()
}
