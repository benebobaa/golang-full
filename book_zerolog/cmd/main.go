package main

import (
	"github.com/benebobaa/hatetepe"
	"github.com/rs/zerolog/log"
)

func main() {

	bookHandler := initHandler()

	router := initRoutes(bookHandler)

	server := hatetepe.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Info().Msg("Starting server on port 8080")

	if err := server.ListenAndServe(); err != nil {
		log.Error().Str("error", err.Error()).Msg("Error when starting the server")
	}
}
