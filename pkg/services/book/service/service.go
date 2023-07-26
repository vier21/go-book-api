package service

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/vier21/go-book-api/pkg/services/book"
	"github.com/vier21/go-book-api/pkg/services/book/common"
	"github.com/vier21/go-book-api/pkg/services/book/model"
)

type BookService struct {
	BookRepository book.BookRepositoryInterface
}

func NewBookService(book book.BookRepositoryInterface) *BookService {
	return &BookService{
		BookRepository: book,
	}
}

func (b *BookService) GetAllBook(ctx context.Context) ([]model.Book, error) {
	book, err := b.BookRepository.FindAll(ctx)

	if err != nil {
		log.Println("error GetAllBook function")
		return book, err
	}

	return book, nil
}

func (b *BookService) GetBookByTitle(ctx context.Context, title string) (model.Book, error) {
	book, err := b.BookRepository.FindByTitle(ctx, title)

	if err != nil {
		return book, err
	}
	return book, nil
}
func (b *BookService) GetBookBySlug(ctx context.Context, slug string) (model.Book, error) {
	book, err := b.BookRepository.FindBySlug(ctx, slug)

	if err != nil {
		return book, err
	}
	return book, nil
}

func (b *BookService) StoreBook(ctx context.Context, books ...model.Book) (common.InsertBookResult, error) {
	var (
		result       common.InsertBookResult
		insertedBook []model.Book
	)

	// this condition is for bulk inserting
	if len(books) > 1 {
		for _, book := range books {
			find, _ := b.BookRepository.FindByTitle(ctx, book.Title)

			if book.Title == find.Title {
				_, err := b.BookRepository.UpdateIncBook(ctx, book.Quantity, find.Id)
				if err != nil {
					log.Println(err)
				}
			} else {
				insertedBook = append(insertedBook, book)
			}
		}

		_, err := b.BookRepository.BulkInsertBook(ctx, insertedBook)

		if err != nil {
			return result, err
		}

		result = common.InsertBookResult{
			Result:      insertedBook,
			ResultCount: len(insertedBook),
		}

		return result, nil
	}

	currBook := books[0]
	prev, _ := b.BookRepository.FindByTitle(ctx, currBook.Title)

	if prev.Title == currBook.Title {

		upd, err := b.BookRepository.UpdateIncBook(ctx, currBook.Quantity, prev.Id)
		if err != nil {
			return result, err
		}

		insertedBook = append(insertedBook, upd)
		result = common.InsertBookResult{
			Result:      insertedBook,
			ResultCount: len(insertedBook),
		}
		
		return result, nil
	}

	book, err := b.BookRepository.InsertBook(ctx, books[0])
	if err != nil {
		return result, err
	}

	insertedBook = append(insertedBook, book)
	result = common.InsertBookResult{
		Result:      insertedBook,
		ResultCount: len(result.Result),
	}

	return result, nil
}

func (b *BookService) UpdateBook(ctx context.Context, id string, book model.Book) (model.Book, error) {
	if reflect.ValueOf(book).IsZero() {
		return model.Book{}, fmt.Errorf("no data in request, your data is up to date")
	}

	if id == "" {
		return model.Book{}, fmt.Errorf("you must specified the id")
	}

	upd, err := b.BookRepository.UpdateBook(ctx, id, book)
	if err != nil {
		log.Println(err)
		return upd, err
	}

	return upd, nil
}

func (b *BookService) DeleteBook(ctx context.Context, id ...string) (int, error) {
	if len(id) > 1 {
		count, err := b.BookRepository.BulkDeleteBook(ctx, id)
		if err != nil {
			return count, err
		}
		return count, nil
	}

	find, err := b.BookRepository.FindById(ctx, id[0])

	if err != nil {
		return 0, err
	}

	if find.Quantity > 1 {
		_, err := b.BookRepository.UpdateIncBook(ctx, -1, find.Id)
		if err != nil {
			return 0, err
		}
		return len(id), nil
	}

	if err := b.BookRepository.DeleteBook(ctx, id[0]); err != nil {
		return 0, err
	}

	return len(id), nil
}
