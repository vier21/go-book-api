package book

import (
	"context"

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

type BookServiceRepository interface {
}
