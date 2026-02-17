package main

import (
	"log"
	"os"

	"net/http"
	"time"

	"github.com/Frank-Macedo/lab-forecast/internal/api/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/temperature/{cep}", handlers.GetTemperature).Methods("GET")
	router.HandleFunc("/", handlers.Welcome).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
