package service

import (
	"github.com/artemsmotritel/uni-architecture-lab4/database"
	"github.com/artemsmotritel/uni-architecture-lab4/types"
	"strings"
)

type AddBookInput struct {
	Title    string
	AuthorId int64
	Type     types.BookType
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
	} else if _, err := v.db.GetUser(input.AuthorId); err != nil {
		v.errors["AuthorId"] = "authorId doesn't exist"
	}

	if strings.TrimSpace(input.Title) == "" {
		v.errors["Title"] = "title is missing"
	}

	if strings.TrimSpace(string(input.Type)) == "" {
		v.errors["Type"] = "type is missing"
	}

	return len(v.errors) == 0
}

func (s *Service) AddBook(input AddBookInput) (types.Book, error) {
	validator := newAddBookValidator(s.DB)

	if ok := validator.validate(input); !ok {
		return types.Book{}, ErrFailedValidation
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
