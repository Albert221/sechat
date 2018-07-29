package main

import (
	"log"
	"net/http"
	"time"

	"github.com/Albert221/sechat/server/api"
	"github.com/gorilla/mux"
	"github.com/Albert221/sechat/server/domain"
	"github.com/Albert221/sechat/server/sql"
)

func main() {
	storage := sql.NewStorage()

	r := mux.NewRouter()

	pool := domain.NewPool(&storage)
	controller := api.NewController(&pool, r)

	r.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte("<form action='/room' method='post'><input type='submit'></form>"))
	}).Methods("GET")

	r.HandleFunc("/room", controller.New).Methods("POST").Name("new_room")
	r.HandleFunc("/room/{id}", controller.Room).Methods("GET").Name("room")

	srv := http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
