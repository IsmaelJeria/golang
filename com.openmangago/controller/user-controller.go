package controller

import (
	"context"
	"log"
	"net/http"

	"com.openmangago/services"

	"cloud.google.com/go/firestore"
)

//Store model..
type Store struct {
	ctx    context.Context
	log    *log.Logger
	client *firestore.Client
}

//NewStore crea una instancia del modelo
func NewStore(ctx context.Context, l *log.Logger, client *firestore.Client) *Store {
	return &Store{
		ctx,
		l,
		client,
	}
}

// GetUsers get all users
func (fm *Store) GetUsers(rw http.ResponseWriter, r *http.Request) {
	usersList, err := services.GetUsers(fm.ctx, fm.client)
	if err != nil {
		fm.log.Printf("%s", err)
	}
	err = usersList.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}
