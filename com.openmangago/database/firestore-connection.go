package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

//CreateClient create client firebase
func CreateClient(ctx context.Context, jsonPath string) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "open-manga-go"

	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}
