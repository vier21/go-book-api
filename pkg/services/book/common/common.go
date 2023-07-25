package common

import "github.com/vier21/go-book-api/pkg/services/book/model"

type InsertBookResponse struct {
	Status      string
	InsertCount string
	Data        interface{}
}

type InsertBookResult struct {
	Result      []model.Book
	ResultCount int
}

type UpdateBookResult struct {
	
}

type IncrementUpdateResult struct {
	TotalInc int
	Book     model.Book
}

type BulkUpdateIncrement struct {
	Data []IncrementUpdateResult
}
