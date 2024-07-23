package handler

import (
	"book_zerolog/internal/domain"
	"book_zerolog/internal/usecase"

	"github.com/benebobaa/hatetepe"
	"github.com/benebobaa/valo"
	"github.com/rs/zerolog/log"
)

type BookHandler struct {
	bc usecase.BookUsecase
}

func NewBookHandler(bc *usecase.BookUsecase) *BookHandler {
	return &BookHandler{
		bc: *bc,
	}
}

func (b *BookHandler) Store(w hatetepe.ResponseWriter, r *hatetepe.Request) {

	var book domain.Book
	var errStr string

	w.SetHeader("Content-Type", "application/json")

	if err := r.ParseJSON(&book); err != nil {
		log.Error().Str("statusCode", "400").Str("error", err.Error()).Msg("Error parsing json book")

		errStr = err.Error()
		w.WriteHeader(400)
		w.WriteJSON(
			domain.BaseResponse[*domain.Book]{
				Error: &errStr,
			},
		)
		return
	}

	err := valo.Validate(book)

	if err != nil {
		log.Error().Str("statusCode", "400").Str("error", err.Error()).Msg("Error validation input book")

		errStr = err.Error()

		w.WriteHeader(400)
		w.WriteJSON(
			domain.BaseResponse[*domain.Book]{
				Error: &errStr,
			},
		)
		return
	}

	b.bc.StoreBook(&book)

	log.Info().Str("statusCode", "201").Interface("data", book).Msg("Success store new book")

	w.WriteHeader(201)
	w.WriteJSON(
		domain.BaseResponse[domain.Book]{
			Data: book,
		},
	)
}

func (b *BookHandler) GetAll(w hatetepe.ResponseWriter, r *hatetepe.Request) {

	books := b.bc.GetAllBooks()

	w.SetHeader("Content-Type", "application/json")

	log.Info().Str("statusCode", "200").Interface("data", books).Msg("Success get all books")

	w.WriteHeader(200)
	w.WriteJSON(
		domain.BaseResponse[[]domain.Book]{
			Data: books,
		},
	)
}
