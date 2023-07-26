package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	var res Response
	find, err := a.Services.GetBookBySlug(r.Context(), slug)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		code := strconv.Itoa(http.StatusBadRequest)
		res = Response{
			Status: fmt.Sprintf("error get data by slug: %s (%s)", err.Error(), code),
			Data:   nil,
		}
		json.NewEncoder(w).Encode(res)
		return
	}

	code := strconv.Itoa(http.StatusOK)
	res = Response{
		Status: fmt.Sprintf("Success (%s)", code),
		Data:   find,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		code := strconv.Itoa(http.StatusInternalServerError)
		res = Response{
			Status: fmt.Sprintf("cannot fetch responses: %s (%s)", err.Error(), code),
			Data:   nil,
		}
		json.NewEncoder(w).Encode(res)
		return
	}
}
func (a *ApiServer) StoreBookHandler(w http.ResponseWriter, r *http.Request) {
	
}
func (a *ApiServer) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {

}
func (a *ApiServer) DeleteBookHandlerHandler(w http.ResponseWriter, r *http.Request) {

}
