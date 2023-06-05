package handler

import (
	"encoding/json"
	"go-library/model"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h Handler) AddAuthor(w http.ResponseWriter, r *http.Request) {
	var author model.Author

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(b, &author)

	err = h.service.AddAuthor(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetAuthor(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(authorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	author, err := h.service.GetAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := h.service.GetAllAuthors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(authors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	authorId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(authorId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}