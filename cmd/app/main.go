package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vier21/go-book-api/pkg/services/user/model"
	"github.com/vier21/go-book-api/pkg/services/user/repository"
	"github.com/vier21/go-book-api/pkg/services/user/service"
)

func main() {
	testup := repository.NewRepository()
	testsvc := service.NewUserService(testup)
	res , err := testsvc.UpdateUser(context.Background(),"8567c0e9-651c-4df0-999e-4ea37d14e7c3", model.UpdateUser{
		Email: "kucing22@gmail.com",
	})
	
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(res)

	// svc := service.NewUserService(repository.NewRepository())
	// defer db.Disconnect()
	// s := server.NewServer(svc)
	// s.Run()
}
