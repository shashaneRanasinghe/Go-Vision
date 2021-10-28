package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/shashaneRanasinghe/Go-Vision/handlers"
)


func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("%v", err)
	}
	handlers.RequestHandler()

}
