package main

import (
	"book_zerolog/internal/handler"
	"book_zerolog/internal/repository"
	"book_zerolog/internal/usecase"

	"github.com/benebobaa/hatetepe"
)

func initHandler() *handler.BookHandler {

	br := repository.NewBookRepository()
	bc := usecase.NewBookUsecase(br)

	return handler.NewBookHandler(bc)
}

func initRoutes(bh *handler.BookHandler) *hatetepe.Router {

	router := hatetepe.NewRouter()

	// Book routes
	router.HandleFunc("POST", "/api/books", bh.Store)
	router.HandleFunc("GET", "/api/books", bh.GetAll)

	return router
}
