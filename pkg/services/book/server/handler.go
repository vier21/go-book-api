package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vier21/go-book-api/pkg/services/book/model"
)

type Response struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func (a *ApiServer) GetAllBookHandler(w http.ResponseWriter, r *http.Request) {
	books, err := a.Services.GetAllBook(r.Context())
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		msg := Response{
			Status: fmt.Sprintf("bad request (%s)", strconv.Itoa(http.StatusBadRequest)),
			Data:   nil,
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

	suc := Response{
		Status: fmt.Sprintf("Success (%s)", strconv.Itoa(http.StatusOK)),
		Data:   books,
	}

	if err := json.NewEncoder(w).Encode(suc); err != nil {
		msg := Response{
			Status: fmt.Sprintf("can't fetch response (%s)", strconv.Itoa(http.StatusInternalServerError)),
			Data:   nil,
		}
		json.NewEncoder(w).Encode(msg)
		return
	}

}
func (a *ApiServer) GetBookBySlugHandler(w http.ResponseWriter, r *http.Request) {
	slug := chi.URLParam(r, "slug")
	find, err := a.Services.GetBookBySlug(r.Context(), slug)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		code := strconv.Itoa(http.StatusBadRequest)

		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("error get data by slug: %s (%s)", err.Error(), code),
			Data:   nil,
		})
		return
	}

	code := strconv.Itoa(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Response{
		Status: fmt.Sprintf("Success (%s)", code),
		Data:   find,
	}); err != nil {
		code := strconv.Itoa(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("cannot fetch responses: %s (%s)", err.Error(), code),
			Data:   nil,
		})
		return
	}
}
func (a *ApiServer) StoreBookHandler(w http.ResponseWriter, r *http.Request) {
	var books []model.Book
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&books); err != nil {
		code := strconv.Itoa(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("error response body : %s (%s)", err.Error(), code),
			Data:   nil,
		})
		return
	}

	store, err := a.Services.StoreBook(r.Context(), books...)

	if err != nil {
		code := strconv.Itoa(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("error storing book: %s (%s)", err.Error(), code),
			Data:   nil,
		})
		return
	}

	code := strconv.Itoa(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Response{
		Status: fmt.Sprintf("Success (%s)", code),
		Data:   store,
	}); err != nil {
		code := strconv.Itoa(http.StatusInternalServerError)

		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("cannot fetch responses: %s (%s)", err.Error(), code),
			Data:   nil,
		})
		return
	}

}
func (a *ApiServer) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book model.Book
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		code := strconv.Itoa(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("error request body not valid: %s (%s)", err.Error(), code),
			Data:   nil,
		})
		return
	}

	upd, err := a.Services.UpdateBook(r.Context(), id, book)

	if err != nil {
		code := strconv.Itoa(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("error updating book: %s (%s)", err.Error(), code),
			Data:   nil,
		})
		return
	}

	err = json.NewEncoder(w).Encode(Response{
		Status: fmt.Sprintf("Success (%s)", strconv.Itoa(http.StatusOK)),
		Data:   upd,
	})

	if err != nil {
		code := strconv.Itoa(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("error updating book: %s (%s)", err.Error(), code),
			Data:   nil,
		})
		return
	}

}

type BulkDelete struct {
	Ids []string `json:"ids"`
}

func (a *ApiServer) DeleteBookHandlerHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Add("Content-Type", "application/json")

	if id == "" {
		var ids []string

		if err := json.NewDecoder(r.Body).Decode(&ids); err != nil {
			code := strconv.Itoa(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Response{
				Status: fmt.Sprintf("error request body not valid: %s (%s)", err.Error(), code),
				Data:   nil,
			})
			return
		}

		count, err := a.Services.DeleteBook(r.Context(), ids...)
		if err != nil {
			code := strconv.Itoa(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{
				Status: fmt.Sprintf("error while deleting item: %s (%s)", err.Error(), code),
				Data:   fmt.Sprintf("deleted items: %s", strconv.Itoa(count)),
			})
			return
		}

		err = json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("success delete items (%s)", strconv.Itoa(http.StatusOK)),
			Data:   fmt.Sprintf("deleted items: %s", strconv.Itoa(count)),
		})

		if err != nil {
			http.Error(w, "error fetching responses", http.StatusInternalServerError)
		}

		return
	}

	count, err := a.Services.DeleteBook(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Status: fmt.Sprintf("error request body not valid: %s (%s)", err.Error(), strconv.Itoa(http.StatusBadRequest)),
			Data:   fmt.Sprintf("deleted items: %s", strconv.Itoa(count)),
		})

		return
	}

	json.NewEncoder(w).Encode(Response{
		Status: fmt.Sprintf("success delete items (%s)", strconv.Itoa(http.StatusOK)),
		Data:   fmt.Sprintf("deleted items: %s", strconv.Itoa(count)),
	})

}
