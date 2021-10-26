package handlers

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/shashaneRanasinghe/Go-Vision/routes"
	"time"

)

func RequestHandler() {
	router := mux.NewRouter()

	server := http.Server{
		Addr:         ":8001",
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	router.HandleFunc("/classify",routes.Classify).
	Methods("POST").
	Headers("content-type", "application/json")

	err := server.ListenAndServe()
	if err != nil{
		log.Fatal(err)
	}
}