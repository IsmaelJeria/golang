package main

import (
	"context"
	"gorillamux/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	sm := mux.NewRouter()
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)

	//CORS
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"http://localhost:4200"}))

	s := &http.Server{
		Addr:         ":8080",
		Handler:      ch(sm),
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
