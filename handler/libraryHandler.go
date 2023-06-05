package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func (h Handler) GetTakenBook(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")

	uid, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	books, err := h.service.GetAllTakenBooksOfUser(uid)
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

func (h Handler) TakeBook(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	bookId := chi.URLParam(r, "book_id")

	uid, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bid, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.TakeBook(uid, bid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) ReturnBook(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	bookId := chi.URLParam(r, "book_id")

	uid, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bid, err := strconv.Atoi(bookId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.PutBookBack(uid, bid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetAllTakenBooks(w http.ResponseWriter, r *http.Request) {
	takenBooks, err := h.service.GetAllTakenBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(takenBooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}

func (h Handler) GetAllTakenBooksOfUser(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	uid, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	takenBooks, err := h.service.GetAllTakenBooksOfUser(uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	b, err := json.Marshal(takenBooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write(b)
	w.WriteHeader(http.StatusOK)
}