package database

import "github.com/artemsmotritel/uni-architecture-lab4/types"

type Database interface {
	GetBooks(filters types.BookFilter) ([]types.Book, error)
	GetBook(id int64) (types.Book, error)
	GetUser(id int64) (types.User, error)
	AddBook(book types.Book) (types.Book, error)
	RemoveBook(id int64) error
	LendBook(id int64) error
	ReturnBook(id int64) error
	AddAuthor(author types.Author) error
	GetAuthors() []types.Author
	GetAuthor(id int64) (types.Author, error)
}
