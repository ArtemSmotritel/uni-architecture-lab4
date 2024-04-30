package database

import "github.com/artemsmotritel/uni-architecture-lab4/types"

type Database interface {
	GetBooks(filters types.BookFilter) ([]types.Book, error)
	GetBook(id int64) (types.Book, error)
	GetUser(id int64) (types.User, error)
	AddBook(book types.Book) (types.Book, error)
	RemoveBook(id int64) error
}
