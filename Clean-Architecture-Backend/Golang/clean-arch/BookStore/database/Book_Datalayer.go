package database

import (
	intf "cleanapp/BookStore/interface"
	"cleanapp/BookStore/model"
	"context"
	"fmt"
)

type BookDatalayerImpl struct {
	books map[string]*model.Book
}

func NewBookDatalayerImpl() intf.BookDatalayer {
	return &BookDatalayerImpl{
		books: make(map[string]*model.Book),
	}
}

func (datalayer *BookDatalayerImpl) TestBookDatalayer(ctx context.Context) {

	fmt.Println("**** Inside Book Datalayer ****")
}

func (datalayer *BookDatalayerImpl) CreateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	id := fmt.Sprintf("%d", len(datalayer.books)+1)
	book.ID = id
	datalayer.books[id] = book
	return book, nil
}

func (datalayer *BookDatalayerImpl) GetBookByID(ctx context.Context, id string) (*model.Book, error) {
	book, exists := datalayer.books[id]
	if !exists {
		return nil, fmt.Errorf("book not found")
	}
	return book, nil
}

func (datalayer *BookDatalayerImpl) DeleteBook(ctx context.Context, id string) error {
	if _, exists := datalayer.books[id]; !exists {
		return fmt.Errorf("book not found")
	}
	delete(datalayer.books, id)
	return nil
}
