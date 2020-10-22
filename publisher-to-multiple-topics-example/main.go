package main

import (
	"context"
	"eventadapter/src/publisher"
	"eventadapter/src/resource"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/buaazp/fasthttprouter"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

func loadEnviromentVariables() {
	godotenv.Load()
}

func startHTTPServer(r *fasthttprouter.Router) {
	log.Fatal(fasthttp.ListenAndServe(":"+os.Getenv("APPLICATION_PORT"), r.Handler))
}

func main() {

	loadEnviromentVariables()
	fmt.Println("application port : " + os.Getenv("APPLICATION_PORT"))
	ctx := context.Background()
	pubSubClient, clientErr := pubsub.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"))
	if clientErr != nil {
		fmt.Println("Client error")
		// Handle error.
	}

	router := fasthttprouter.New()

	asnVerifyTopic := pubSubClient.Topic(os.Getenv("ASN_VERIFY_TOPIC_NAME"))
	asnVerifyPublisher := publisher.NewAsnVerifyPublisher(ctx, asnVerifyTopic)
	resource.NewAsnVerifyResource(asnVerifyPublisher, router)

	distributionOrderUpdateTopic := pubSubClient.Topic(os.Getenv("DISTRIBUTION_ORDER_UPDATE_TOPIC_NAME"))
	distributionOrderUpdatePublisher := publisher.NewDistributionOrderUpdatePublisher(ctx, distributionOrderUpdateTopic)
	resource.NewDistributionOrderUpdateResource(distributionOrderUpdatePublisher, router)

	inventoryEventTopic := pubSubClient.Topic(os.Getenv("INVENTORY_EVENT_TOPIC_NAME"))
	inventoryEventPublisher := publisher.NewInventoryEventPublisher(ctx, inventoryEventTopic)
	resource.NewInventoryEventResource(inventoryEventPublisher, router)

	inventorySyncTopic := pubSubClient.Topic(os.Getenv("INVENTORY_SYNC_TOPIC_NAME"))
	inventorySyncPublisher := publisher.NewInventorySyncPublisher(ctx, inventorySyncTopic)
	resource.NewInventorySyncResource(inventorySyncPublisher, router)

	doCancelTopic := pubSubClient.Topic(os.Getenv("DISTRIBUTION_ORDER_CANCELLATION_TOPIC_NAME"))
	doCancelPublisher := publisher.NewDOCancelPublisher(ctx, doCancelTopic)
	resource.NewDOCancelResource(doCancelPublisher, router)

	asnShippingTopic := pubSubClient.Topic(os.Getenv("ASN_SHIPPING_TOPIC_NAME"))
	asnShippingPublisher := publisher.NewASNShippingPublisher(ctx, asnShippingTopic)
	resource.NewASNShippingResource(asnShippingPublisher, router)

	startHTTPServer(router)
}
