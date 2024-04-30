package service

import (
	"fmt"
	"github.com/artemsmotritel/uni-architecture-lab4/database"
	"github.com/artemsmotritel/uni-architecture-lab4/types"
	"strings"
)

type AddBookInput struct {
	Title    string         `json:"title"`
	AuthorId int64          `json:"authorId"`
	Type     types.BookType `json:"type"`
}

type addBookValidator struct {
	errors map[string]string
	db     database.Database
}

func newAddBookValidator(db database.Database) *addBookValidator {
	return &addBookValidator{
		errors: make(map[string]string),
		db:     db,
	}
}

func (v *addBookValidator) validate(input AddBookInput) bool {
	if input.AuthorId == 0 {
		v.errors["AuthorId"] = "authorId is missing"
	} else if _, err := v.db.GetAuthor(input.AuthorId); err != nil {
		v.errors["AuthorId"] = "authorId doesn't exist"
	}

	if strings.TrimSpace(input.Title) == "" {
		v.errors["Title"] = "title is missing"
	}

	if bookType := strings.TrimSpace(string(input.Type)); bookType == "" {
		v.errors["Type"] = "type is missing"
	} else if bookType != string(types.PAPER_BOOK) && bookType != string(types.E_BOOK) {
		v.errors["Type"] = "type is invalid"
	}

	return len(v.errors) == 0
}

func (s *Service) AddBook(input AddBookInput) (types.Book, error) {
	validator := newAddBookValidator(s.DB)

	if ok := validator.validate(input); !ok {
		values := make([]string, 0, len(validator.errors))

		for _, err := range validator.errors {
			values = append(values, err)
		}

		return types.Book{}, fmt.Errorf("%s, %w", strings.Join(values, ";"), ErrFailedValidation)
	}

	book := types.NewBook(0, input.Title, input.AuthorId, input.Type, types.AVAILABLE)

	savedBook, err := s.DB.AddBook(book)
	if err != nil {
		return book, err
	}

	return savedBook, nil
}

func (s *Service) GetBooks(statuses []string) ([]types.Book, error) {
	var filter types.BookFilter

	for _, status := range statuses {
		switch status {
		case string(types.AVAILABLE):
			filter.Statuses = append(filter.Statuses, types.AVAILABLE)
		case string(types.LENT):
			filter.Statuses = append(filter.Statuses, types.LENT)
		case string(types.LOST):
			filter.Statuses = append(filter.Statuses, types.LOST)
		case string(types.OVERDUE):
			filter.Statuses = append(filter.Statuses, types.OVERDUE)
		}
	}

	return s.DB.GetBooks(filter)
}

func (s *Service) RemoveBook(id int64) error {
	return s.DB.RemoveBook(id)
}

func (s *Service) GetBook(id int64) (types.Book, error) {
	return s.DB.GetBook(id)
}

func (s *Service) LendBook(id int64) error {
	return s.DB.LendBook(id)
}

func (s *Service) ReturnBook(id int64) error {
	return s.DB.ReturnBook(id)
}
