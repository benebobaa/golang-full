package main

import (
	"log"
	"net/http"
)

func main() {

	eventHandler, ticketHandler, orderHandler := initHandler()

	router := initRouter(eventHandler, ticketHandler, orderHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}
