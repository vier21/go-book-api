package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/vier21/go-book-api/config"
	"github.com/vier21/go-book-api/pkg/services/user"
	"github.com/vier21/go-book-api/pkg/services/user/middleware"
)

type ApiServer struct {
	Server  *http.Server
	Mux     *http.ServeMux
	Service user.UserService
}

func NewServer(svc user.UserService) *ApiServer {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:         config.GetConfig().ServerPort,
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	return &ApiServer{
		Server:  server,
		Mux:     mux,
		Service: svc,
	}
}

func (a *ApiServer) Run() error {
	mux := a.Mux

	mux.HandleFunc("/api/v1/register", middleware.TimoutMiddleware(a.RegisterHandler))
	mux.HandleFunc("/api/v1/login", a.LoginHandler)
	mux.HandleFunc("/api/v1/user", middleware.VerifyJWT(a.GetUserHandler))
	
	go func() {
		log.Printf("Server start on localhost%s \n", config.GetConfig().ServerPort)
		err := a.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Error: %s \n", err)
		}
	}()

	a.GracefullShutdown()
	return nil
}

func (a *ApiServer) GracefullShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped gracefully")
}
