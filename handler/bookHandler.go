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
	author_id := chi.URLParam(r, "author_id")
	bookId := chi.URLParam(r, "id")

	bid, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aid, err := strconv.Atoi(author_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	book, err := h.service.GetBook(bid, aid)
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
	author_id := chi.URLParam(r, "author_id")
	bookId := chi.URLParam(r, "id")

	bid, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	aid, err := strconv.Atoi(author_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteBook(bid, aid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetAllAuthorBooks(w http.ResponseWriter, r *http.Request) {
	author_id := chi.URLParam(r, "author_id")

	aid, err := strconv.Atoi(author_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	books, err := h.service.GetAllAuthorBooks(aid)
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