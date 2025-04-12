package intf

import (
	"cleanapp/BookStore/model"
	"context"
)

type BookService interface {
	PrintBookTitle(ctx context.Context, book *model.Book)
	TestBookService(ctx context.Context)
	CreateBook(ctx context.Context, book *model.Book) (*model.Book, error)
	GetBookByID(ctx context.Context, id string) (*model.Book, error)
	DeleteBook(ctx context.Context, id string) error
}

type BookDatalayer interface {
	TestBookDatalayer(ctx context.Context)
	CreateBook(ctx context.Context, book *model.Book) (*model.Book, error)
	GetBookByID(ctx context.Context, id string) (*model.Book, error)
	DeleteBook(ctx context.Context, id string) error
}
