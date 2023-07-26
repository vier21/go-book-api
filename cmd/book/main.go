package main

import (
	"github.com/vier21/go-book-api/pkg/services/book/repository"
	"github.com/vier21/go-book-api/pkg/services/book/server"
	"github.com/vier21/go-book-api/pkg/services/book/service"
)

func main() {
	repo := repository.NewBookRepository()
	svc := service.NewBookService(repo)

	s := server.NewServer(svc)
	s.Run()
}
