package handler

import (
	"encoding/json"
	"go-library/model"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h Handler) AddBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.Unmarshal(b, &book)

	err = h.service.AddBook(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := h.service.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetAllBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := chi.URLParam(r, "id")

	id, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}