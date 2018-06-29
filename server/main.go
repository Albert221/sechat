package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Albert221/sechat/server/api"
	"github.com/gorilla/mux"
)

func main() {
	controller := api.NewController()

	r := mux.NewRouter()
	r.HandleFunc("/new", controller.NewEndpoint).Methods("POST")
	r.HandleFunc("/{id}", controller.ChatEndpoint).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}