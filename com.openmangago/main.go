package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"com.openmangago/controller"
	"com.openmangago/database"
	"com.openmangago/env"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	//log to console
	l := log.New(os.Stdout, "openmangago-api ", log.LstdFlags)

	//Enviroment path
	jsonPath := env.GoDotEnvVariable("CLOUD_PATH")
	l.Printf("godotenv : %s = %s \n", "CLOUD_PATH", jsonPath)
	// Get a Firestore client.
	ctx := context.Background()
	client := database.CreateClient(ctx, jsonPath)
	fm := controller.NewStore(ctx, l, client)
	defer client.Close()

	//router
	r := mux.NewRouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/users", fm.GetUsers)

	//CORS
	ch := handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:4200"}))

	//server
	s := &http.Server{
		Addr:         ":8080",
		Handler:      ch(r),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		l.Println("Starting server on port 8080")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("error starting server: %s\n", err)
		}
	}()
	// trap sigterm or interupt and gracefully shutdown the server
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	// Block until a signal is received.
	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	d := time.Now().Add(30 * time.Second)
	tc, cancel := context.WithDeadline(context.Background(), d)
	cancel()
	s.Shutdown(tc)
}
