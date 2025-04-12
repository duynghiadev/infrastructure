package service

import (
	intf "cleanapp/BookStore/interface"
	"cleanapp/BookStore/model"
	"context"
	"fmt"
)

type BookServiceImpl struct {
	BookDatalayer intf.BookDatalayer
}

// This function will create an BookService object which represents the BookService interface
func NewBookServiceImpl(bookDataLayer intf.BookDatalayer) intf.BookService {

	return &BookServiceImpl{
		BookDatalayer: bookDataLayer,
	}

}

func (service *BookServiceImpl) PrintBookTitle(ctx context.Context, book *model.Book) {

}

func (service *BookServiceImpl) TestBookService(ctx context.Context) {

	fmt.Println("**** Inside Book Service ****")

	service.BookDatalayer.TestBookDatalayer(ctx)

}

func (service *BookServiceImpl) CreateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	return service.BookDatalayer.CreateBook(ctx, book)
}

func (service *BookServiceImpl) GetBookByID(ctx context.Context, id string) (*model.Book, error) {
	return service.BookDatalayer.GetBookByID(ctx, id)
}

func (service *BookServiceImpl) DeleteBook(ctx context.Context, id string) error {
	return service.BookDatalayer.DeleteBook(ctx, id)
}

func (service *BookServiceImpl) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	return service.BookDatalayer.GetAllBooks(ctx)
}
