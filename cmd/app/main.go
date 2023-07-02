package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/google/uuid"
	"github.com/vier21/go-book-api/pkg/db"
	"github.com/vier21/go-book-api/pkg/services/user/model"
	"github.com/vier21/go-book-api/pkg/services/user/repository"
)

type Server struct {
	Port string
}

func Any(id ...string) {
	fmt.Println(len(id))

}

func main() {
	db := db.NewConnection()
	defer db.Disconnect()

	db.Ping()

	repo := repository.NewRepository(db)

	result, err := repo.UpdateUser(context.TODO(), model.User{
		Id:       "bd6d4fdsds",
		Username: "kuntul",
		Password: "babak210901",
		Email:    "XavierSamuel@gmail.com",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result, err)

	Any("marco")

	if err := repo.DeleteUser(context.TODO(), "bd6d4f36-9cd7-420f-8ae2-a9343f92f6dd"); err != nil {
		fmt.Println(err)
	}
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	rootApp := strings.TrimSuffix(path, "/bin/config")

	fmt.Println(rootApp)
	os.Setenv("APP_PATH", rootApp)

	fmt.Println(os.Getenv("APP_PATH"))
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
