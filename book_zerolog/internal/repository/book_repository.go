package repository

import (
	"book_zerolog/internal/domain"
	"log"
)

type BookRepository struct {
	books  map[int]domain.Book
	lastID int
}

func NewBookRepository() *BookRepository {
	return &BookRepository{
		books:  make(map[int]domain.Book),
		lastID: 0,
	}
}

func (b *BookRepository) Save(value *domain.Book) {

	b.lastID += 1
	log.Println("last id :: ", b.lastID)
	value.ID = b.lastID

	b.books[value.ID] = *value
}

func (b *BookRepository) FindAll() []domain.Book {
	var books []domain.Book

	for _, v := range b.books {
		books = append(books, v)
	}

	return books
}
