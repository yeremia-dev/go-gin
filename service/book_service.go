package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/yeremia-dev/go-gin/dto"
	"github.com/yeremia-dev/go-gin/entity"
	"github.com/yeremia-dev/go-gin/repository"
)

type BookService interface {
	Insert(book dto.BookCreateDTO) entity.Book
	Update(book dto.BookUpdateDTO) entity.Book
	Delete(book entity.Book)
	All() []entity.Book
	FindByID(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	rep repository.BookRepository
}

func NewBookService(rep repository.BookRepository) BookService {
	return &bookService{
		rep: rep,
	}
}

func (service *bookService) Insert(book dto.BookCreateDTO) entity.Book {

	bookToInsert := entity.Book{}
	err := smapping.FillStruct(&bookToInsert, smapping.MapFields(&book))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.rep.InsertBook(bookToInsert)
	return res

}

func (service *bookService) Update(book dto.BookUpdateDTO) entity.Book {

	bookToUpdate := entity.Book{}
	err := smapping.FillStruct(&bookToUpdate, smapping.MapFields(&book))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := service.rep.UpdateBook(bookToUpdate)
	return res

}

func (service *bookService) Delete(book entity.Book) {

	service.rep.DeleteBook(book)

}

func (service *bookService) All() []entity.Book {

	return service.rep.AllBook()

}

func (service *bookService) FindByID(bookID uint64) entity.Book {

	return service.rep.FindBookById(bookID)

}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {

	book := service.rep.FindBookById(bookID)
	id := fmt.Sprintf("%v", book.UserId)
	return userID == id

}
