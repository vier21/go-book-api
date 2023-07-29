package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vier21/go-book-api/config"
	mw "github.com/vier21/go-book-api/pkg/middleware"
	"github.com/vier21/go-book-api/pkg/services/user"
)

type ApiServer struct {
	Server  *http.Server
	Router  *chi.Mux
	Service user.UserService
}

func NewServer(svc user.UserService) *ApiServer {
	router := chi.NewRouter()

	server := &http.Server{
		Addr:         config.GetConfig().ServerPort,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	return &ApiServer{
		Server:  server,
		Router:  router,
		Service: svc,
	}
}

func (a *ApiServer) NewRouter() *chi.Mux {
	return a.Router
}

func (a *ApiServer) Run() error {
	r := a.NewRouter()

	r.Route("/api/v1/", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Group(func(r chi.Router) {
			r.Use(middleware.Timeout(10 * time.Second))
			r.Use(mw.VerifyJWT)

			r.Get("/user", a.GetCurrentUserHandler)
			r.Put("/user/{id}", a.UpdateUserHandler)
			r.Delete("/user/{id}", a.DeleteUserHandler)
			r.Delete("/user", a.DeleteUserHandler)
		})

		r.Group(func(r chi.Router) {
			r.Use(mw.TimoutMiddleware)
			r.Post("/register", a.RegisterHandler)
			r.Post("/login", a.LoginHandler)
		})
	})

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
