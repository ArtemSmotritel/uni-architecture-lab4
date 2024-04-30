package types

import "errors"

var (
	ErrBookNotExist       = errors.New("book does not exist")
	ErrBookStatusConflict = errors.New("book status conflict")
)

type BookType string

const (
	E_BOOK     BookType = "E_BOOK"
	PAPER_BOOK BookType = "PAPER_BOOK"
)

type BookStatus string

const (
	AVAILABLE BookStatus = "AVAILABLE"
	LENT      BookStatus = "LENT"
	LOST      BookStatus = "LOST"
	OVERDUE   BookStatus = "OVERDUE"
)

type Book struct {
	ID       int64      `json:"id"`
	Title    string     `json:"title"`
	AuthorId int64      `json:"authorId"`
	Type     BookType   `json:"type"`
	Status   BookStatus `json:"status"`
}

type BookFilter struct {
	Statuses []BookStatus
}

func NewBook(id int64, title string, authorId int64, bookType BookType, status BookStatus) Book {
	return Book{
		ID:       id,
		Title:    title,
		AuthorId: authorId,
		Type:     bookType,
		Status:   status,
	}
}
