package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"publisher-archetype/src/publisher"
	"publisher-archetype/src/resource"

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
	fmt.Println("application port: " + os.Getenv("APPLICATION_PORT"))
	fmt.Println("id proyect: " + os.Getenv("GOOGLE_PROJECT_ID"))
	fmt.Println("topic name: " + os.Getenv("EXAMPLE_TOPIC_NAME"))
	ctx := context.Background()
	//add publisher client here
	pubSubClient, clientErr := pubsub.NewClient(ctx, os.Getenv("GOOGLE_PROJECT_ID"))
	if clientErr != nil {
		fmt.Println("Client error", clientErr)
		// Handle error.
	}

	router := fasthttprouter.New()
	//add topics here
	exampleTopic := pubSubClient.Topic(os.Getenv("EXAMPLE_TOPIC_NAME"))
	examplePublisher := publisher.NewExamplePublisher(ctx, exampleTopic)
	resource.NewExampleResource(examplePublisher, router)

	startHTTPServer(router)
}
