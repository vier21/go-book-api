package main

import (
	"github.com/vier21/go-book-api/pkg/db"
	"github.com/vier21/go-book-api/pkg/services/user/repository"
	"github.com/vier21/go-book-api/pkg/services/user/server"
	"github.com/vier21/go-book-api/pkg/services/user/service"
)

func main() {
	svc := service.NewUserService(repository.NewRepository())
	defer db.Disconnect()
	s := server.NewServer(svc)
	s.Run()
}
