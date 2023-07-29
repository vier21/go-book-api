package main

import (
	"github.com/vier21/go-book-api/pkg/services/user/server"
	"github.com/vier21/go-book-api/pkg/services/user/repository"
	"github.com/vier21/go-book-api/pkg/services/user/service"
)

func main() {
	repo := repository.NewRepository()
	svc := service.NewUserService(repo)
	s := server.NewServer(svc)
	s.Run()
}
