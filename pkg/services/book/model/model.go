package model

import (
	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID `json="id, omitempty"`
	Title     string    `json="title"`
	Author    string    `json="author"`
	Slug      string    `json="slug"`
	Body      string    `json="string"`
	Publisher string    `json="publisher"`
}