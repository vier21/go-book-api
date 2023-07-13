package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vier21/go-book-api/pkg/services/user/common"
	"github.com/vier21/go-book-api/pkg/services/user/model"
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

	result := common.RegisterResponse{
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
	var req common.LoginRequest
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
	resp := common.LoginResponse{
		Status:  status,
		Token:   token,
		Payload: loginpayload,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, ErrFetchResp, http.StatusInternalServerError)
		return
	}
}

func (a *ApiServer) GetCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, ErrMethodNotAllow, http.StatusBadRequest)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(r.Context().Value("data")); err != nil {
		http.Error(w, ErrFetchResp, http.StatusInternalServerError)
		return
	}
}

func (a *ApiServer) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var data model.UpdateUser

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, ErrReqBodyNotValid, http.StatusBadRequest)
		return
	}

	doc, err := a.Service.UpdateUser(r.Context(), id, data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	httpcode := strconv.Itoa(http.StatusOK)
	status := fmt.Sprintf("Success updated user (%s)", httpcode)

	resp := common.UpdateResponse{
		Status:  status,
		Payload: doc,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type BulkDelete struct {
	DeletedID []string `json:"deleteId"`
}

func (a *ApiServer) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Add("Content-Type", "application/json")

	resp := common.DeleteResponse{
		Status: fmt.Sprintf("success %s: ", strconv.Itoa(http.StatusOK)),
		Message: "delete user successfull",	
	}

	if id == "" {
		var data BulkDelete
		
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = a.Service.BulkDeleteUser(r.Context(), data.DeletedID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		byte, err := json.Marshal(resp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(byte)
		return
	}

	err := a.Service.DeleteUser(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
