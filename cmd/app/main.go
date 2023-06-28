package main

import (
	"context"
	"fmt"
	"net/http"

	_ "github.com/google/uuid"
	"github.com/vier21/go-book-api/pkg/db"
	"github.com/vier21/go-book-api/pkg/services/User/model"
	"github.com/vier21/go-book-api/pkg/services/User/repository"
)

type Server struct {
	Port string
}

func Any(id ...string)  {
	fmt.Println(len(id))

}

func main() {
	db := db.NewConnection()
	defer db.Disconnect()

	db.Ping()

	repo := repository.NewRepository(db)

	result, err := repo.UpdateUser(context.TODO(), model.User{
		Id: "bd6d4f",
		Username: "kuntul",
		Password: "babak210901",
		Email: "XavierSamuel@gmail.com",
	})
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result, err)

	Any("marco")

	if err := repo.DeleteUser(context.TODO(), "bd6d4f36-9cd7-420f-8ae2-a9343f92f6dd") ; err != nil {
		fmt.Println(err)
	}
}

func NewServer(port string) *Server {
	return &Server{
		Port: port,
	}
}

func (s *Server) Run() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Project %s", r.Host)
		fmt.Println("request from /")
	})

	err := s.StartServer()

	if err != nil {
		fmt.Println(err)
	}

}

func (s *Server) StartServer() error {
	fmt.Println("Serving on localhost:8000")
	err := http.ListenAndServe(s.Port, nil)

	if err != nil {
		return err
	}
	return nil
}
