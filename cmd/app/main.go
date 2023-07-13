package main

import (
	"context"
	"fmt"

	"github.com/vier21/go-book-api/pkg/services/book/repository"
)

func main() {
	rep := repository.NewBookRepository()
	del := []string{"5fa87425-080a-4992-910c-ba97b35bd60c", "d830af41-5a15-4476-8f22-5bae39a84bae", "f556c8bf-d40a-4990-99fb-e98d19cfc982"}
	_, err := rep.BulkDeleteBook(context.Background(), del)

	if err != nil {
		fmt.Println(err)
		return
	}

	// svc := service.NewUserService(repository.NewRepository())
	// defer db.Disconnect()
	// s := server.NewServer(svc)
	// s.Run()

}
