package main

import (
	"github.com/vier21/go-book-api/pkg/db"
	"github.com/vier21/go-book-api/pkg/services/user/repository"
	"github.com/vier21/go-book-api/pkg/services/user/server"

	"github.com/vier21/go-book-api/pkg/services/user/service"
)

type Server struct {
	Port string
}

func main() {
	defer db.Disconnect()

	svc := service.NewUserService(repository.NewRepository())

	s := server.NewServer(svc)
	s.Run()
}
