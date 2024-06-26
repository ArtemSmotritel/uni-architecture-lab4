package api

import (
	"encoding/json"
	"errors"
	"github.com/artemsmotritel/uni-architecture-lab4/service"
	"github.com/artemsmotritel/uni-architecture-lab4/types"
	"net/http"
	"strconv"
)

func (s *LibraryServer) handleAddBook(w http.ResponseWriter, r *http.Request) {
	var input service.AddBookInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := s.Service.AddBook(input)
	if err != nil {
		if errors.Is(err, service.ErrFailedValidation) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "Something went terribly wrong", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "Something went terribly wrong", http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (s *LibraryServer) handleGetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := s.Service.GetBooks(r.URL.Query()["status"])
	if err != nil {
		http.Error(w, "Something went terribly wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, "Something went terribly wrong", http.StatusInternalServerError)
	}
}

func (s *LibraryServer) handleRemoveBook(w http.ResponseWriter, r *http.Request) {
	bookIdStr := r.PathValue("id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Wrong book id format", http.StatusBadRequest)
		return
	}

	if err := s.Service.RemoveBook(bookId); err != nil {
		if errors.Is(err, types.ErrBookNotExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Something went terribly wrong", http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}

func (s *LibraryServer) handleGetBook(w http.ResponseWriter, r *http.Request) {
	bookIdStr := r.PathValue("id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Wrong book id format", http.StatusBadRequest)
		return
	}

	book, err := s.Service.GetBook(bookId)
	if err != nil {
		if errors.Is(err, types.ErrBookNotExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, "Something went terribly wrong", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, "Something went terribly wrong", http.StatusInternalServerError)
	}
}

func (s *LibraryServer) handleLendBook(w http.ResponseWriter, r *http.Request) {
	bookIdStr := r.PathValue("id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Wrong book id format", http.StatusBadRequest)
		return
	}

	if err := s.Service.LendBook(bookId); err != nil {
		if errors.Is(err, types.ErrBookNotExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if errors.Is(err, types.ErrBookStatusConflict) {
			http.Error(w, err.Error(), http.StatusConflict)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *LibraryServer) handleReturnBook(w http.ResponseWriter, r *http.Request) {
	bookIdStr := r.PathValue("id")
	bookId, err := strconv.ParseInt(bookIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Wrong book id format", http.StatusBadRequest)
		return
	}

	if err := s.Service.ReturnBook(bookId); err != nil {
		if errors.Is(err, types.ErrBookNotExist) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else if errors.Is(err, types.ErrBookStatusConflict) {
			http.Error(w, err.Error(), http.StatusConflict)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *LibraryServer) handleGetAuthors(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s.Service.GetAuthors())
}
