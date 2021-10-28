package handlers

import (
	"github.com/gorilla/mux"
	"github.com/shashaneRanasinghe/Go-Vision/routes"
	"log"
	"net/http"
	"time"
)

//the RequestHandler function creates the router and handles the requests
func RequestHandler() {
	router := mux.NewRouter()

	server := http.Server{
		Addr:         ":8001",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	router.HandleFunc("/classify", routes.Classify).
		Methods("POST").
		Headers("content-type", "application/json")

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
