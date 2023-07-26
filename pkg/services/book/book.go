package book

import (
	"context"

	"github.com/vier21/go-book-api/pkg/services/book/common"
	"github.com/vier21/go-book-api/pkg/services/book/model"
)

type BookRepositoryInterface interface {
	FindById(context.Context, string) (model.Book, error)
	FindByTitle(context.Context, string) (model.Book, error)
	FindBySlug(context.Context, string) (model.Book, error)
	FindAll(context.Context) ([]model.Book, error)
	InsertBook(context.Context, model.Book) (model.Book, error)
	BulkInsertBook(context.Context, []model.Book) ([]interface{}, error)
	DeleteBook(context.Context, string) error
	BulkDeleteBook(context.Context, []string) (int, error)
	UpdateBook(context.Context, string, model.Book) (model.Book, error)
	UpdateIncBook(context.Context, int, string) (model.Book, error)
}

type BookServiceInterface interface {
	GetAllBook(context.Context) ([]model.Book, error)
	GetBookByTitle(context.Context, string) (model.Book, error)
	GetBookBySlug(context.Context, string) (model.Book, error)
	StoreBook(context.Context, ...model.Book) (common.InsertBookResult, error)
	UpdateBook(context.Context, string, model.Book) (model.Book, error)
	DeleteBook(context.Context, ...string) (int, error)
}
