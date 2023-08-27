package main

import (
	"gvozdevLabApi/internal/database"
	"gvozdevLabApi/internal/service"
	"gvozdevLabApi/internal/transport/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	base, err := database.Connect()
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	repoDB := database.NewRepoDB(base)
	services := service.NewService(repoDB)
	handlers := handlers.NewHandler(services)
	handlers.GvozdevLabRoutes(r)

	http.Handle("/", r)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}
