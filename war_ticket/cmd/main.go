package main

import (
	"log"
	"net/http"
)

func main() {

	eventHandler, ticketHandler := initHandler()

	router := initRouter(eventHandler, ticketHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}
