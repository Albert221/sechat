package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Albert221/sechat/server/api"
	"github.com/gorilla/mux"
)

func main() {
	repository := api.NewInMemoryChatRepository()
	controller := api.NewController(
		&repository)

	// TODO: Run goroutine to remove rooms older than 24hr

	r := mux.NewRouter()
	r.HandleFunc("/new", controller.NewEndpoint).Methods("POST")
	// Put public key in Authorization header, like this: `Authorization: Key <public key>`
	r.HandleFunc("/{id}", controller.ChatEndpoint).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}