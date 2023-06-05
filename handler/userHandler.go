package handler

import (
	"encoding/json"
	"go-library/model"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(b, &user)

	err = h.service.AddUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}