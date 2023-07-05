package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/vier21/go-book-api/pkg/services/user/model"
	"github.com/vier21/go-book-api/utils"
)

var (
	ErrFetchResp       = "fail to fetch responses"
	ErrMethodNotAllow  = "method not allowed"
	ErrReqBodyNotValid = "request body not valid"
)

func (a *ApiServer) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, ErrMethodNotAllow, http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	var req model.User
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, ErrReqBodyNotValid, http.StatusBadRequest)
		return
	}

	register, err := a.Service.RegisterUser(r.Context(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpcode := strconv.Itoa(http.StatusOK)
	status := fmt.Sprintf("Success (%s)", httpcode)

	result := utils.RegisterResponse{
		Status:  status,
		Payload: register,
	}

	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "fail to fetch responses", http.StatusInternalServerError)
		return
	}

}

func (a *ApiServer) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, ErrMethodNotAllow, http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	var req utils.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, ErrReqBodyNotValid, http.StatusBadRequest)
		return
	}

	loginpayload, token, err := a.Service.LoginUser(r.Context(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	httpcode := strconv.Itoa(http.StatusOK)
	status := fmt.Sprintf("Success (%s)", httpcode)

	resp := utils.LoginResponse{
		Status:  status,
		Token:   token,
		Payload: loginpayload,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, ErrFetchResp, http.StatusInternalServerError)
	}
}

func (a *ApiServer) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrMethodNotAllow, http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello from user %s", r.URL.Path)
}

func (a *ApiServer) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (a *ApiServer) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

}
