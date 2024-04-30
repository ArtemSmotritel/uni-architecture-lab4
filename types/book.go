package types

import "errors"

var (
	ErrBookNotExist = errors.New("book does not exist")
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
	ID       int64
	Title    string
	AuthorId int64
	Type     BookType
	Status   BookStatus
}

type BookFilter struct {
	Statuses []BookStatus
}

func CopyBook(book Book) Book {
	return Book{
		ID:       book.ID,
		Title:    book.Title,
		AuthorId: book.AuthorId,
		Type:     book.Type,
		Status:   book.Status,
	}
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
