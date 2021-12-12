package repository

import (
	"github.com/yeremia-dev/go-gin/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	InsertBook(book entity.Book) entity.Book
	UpdateBook(book entity.Book) entity.Book
	DeleteBook(book entity.Book)
	AllBook() []entity.Book
	FindBookById(bookID uint64) entity.Book
}

type bookRepository struct {
	connection *gorm.DB
}

func NewBookRepository(dbConn *gorm.DB) BookRepository {
	return &bookRepository{
		connection: dbConn,
	}
}

func (rep *bookRepository) InsertBook(book entity.Book) entity.Book {

	rep.connection.Save(&book)
	rep.connection.Preload("User").Find(&book)
	return book

}

func (rep *bookRepository) UpdateBook(book entity.Book) entity.Book {
	rep.connection.Save(&book)
	rep.connection.Preload("User").Find(&book)
	return book
}

func (rep *bookRepository) DeleteBook(book entity.Book) {
	rep.connection.Delete(&book)
}

func (rep *bookRepository) FindBookById(bookID uint64) entity.Book {
	var book entity.Book
	rep.connection.Preload("User").Find(&book, bookID)
	return book
}

func (rep *bookRepository) AllBook() []entity.Book {
	var books []entity.Book
	rep.connection.Preload("User").Find(&books)
	return books
}
