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
	"github.com/vier21/go-book-api/pkg/services/book"
)

type ApiServer struct {
	Services book.BookServiceInterface
	Router   *chi.Mux
	Server   *http.Server
}

func NewServer(booksvc book.BookServiceInterface) *ApiServer {
	mux := chi.NewRouter()

	return &ApiServer{
		Services: booksvc,
		Router:   mux,
		Server: &http.Server{
			Addr:         config.GetConfig().ServerPort,
			Handler:      mux,
			IdleTimeout:  120 * time.Second,
			WriteTimeout: 1 * time.Second,
			ReadTimeout:  1 * time.Second,
		},
	}
}

func (a *ApiServer) NewRouter() *chi.Mux {
	return a.Router
}

func (a *ApiServer) Run() {
	r := a.NewRouter()

	r.Use(middleware.AllowContentType("application/json", "text/xml"))
	r.Route("/api/v1", func(r chi.Router) {
		r.Use(mw.VerifyJWT)
		r.Get("/book", a.GetAllBookHandler)
		r.Get("/{slug}/book", a.GetBookBySlugHandler)
		r.Post("/book", a.StoreBookHandler)
		r.Put("/{id}/book", a.UpdateBookHandler)
		r.Delete("/book/{id}", a.DeleteBookHandlerHandler)
		r.Delete("/book/", a.DeleteBookHandlerHandler)
	})

	go func() {
		log.Printf("Server start on localhost%s \n", config.GetConfig().ServerPort)
		err := a.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server Error: %s \n", err)
		}
	}()

	a.GracefullShutdown()

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
