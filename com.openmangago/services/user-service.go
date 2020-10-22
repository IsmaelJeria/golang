package services

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"com.openmangago/model"
	"google.golang.org/api/iterator"
)

//CreateUserAccount crea cuentas de usuario
func CreateUserAccount(ctx context.Context, client *firestore.Client, user *model.User) error {
	_, _, err := client.Collection("users").Add(ctx, user)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
		return err
	}
	return nil
}

//GetUsers obtiene todos los user
func GetUsers(ctx context.Context, client *firestore.Client) (model.Users, error) {
	var usersList model.Users
	var temp model.User

	iter := client.Collection("users").Documents(ctx)
	i := 0
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		doc.DataTo(&temp)
		usersList = append(usersList, temp)
		i++
	}
	return usersList, nil
}
