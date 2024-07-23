package usecase

import (
	"book_zerolog/internal/domain"
	"book_zerolog/internal/repository"
)

type BookUsecase struct {
	br repository.BookRepository
}

func NewBookUsecase(br *repository.BookRepository) *BookUsecase {
	return &BookUsecase{
		br: *br,
	}
}

func (b *BookUsecase) StoreBook(request *domain.Book) *domain.Book {
	b.br.Save(request)

	return request
}

func (b *BookUsecase) GetAllBooks() []domain.Book {
	return b.br.FindAll()
}
