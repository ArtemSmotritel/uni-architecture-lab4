package database

import (
	"fmt"
	"github.com/artemsmotritel/uni-architecture-lab4/types"
)

type InMemoryDb struct {
	books   []types.Book
	authors []types.Author
	users   []types.User
}

var (
	bookId   int64 = 0
	authorId int64 = 0
	userId   int64 = 0
)

func NewInMemoryDb() *InMemoryDb {
	return &InMemoryDb{
		books:   make([]types.Book, 0),
		authors: make([]types.Author, 0),
	}
}

func (i *InMemoryDb) GetBooks(filters types.BookFilter) ([]types.Book, error) {
	res := make([]types.Book, 0)

	for _, book := range i.books {
		res = append(res, types.CopyBook(book))
	}

	return res, nil
}

func (i *InMemoryDb) GetBook(id int64) (types.Book, error) {
	for _, b := range i.books {
		if b.ID == id {
			return types.CopyBook(b), nil
		}
	}

	return types.Book{}, fmt.Errorf("get book with id %d not found;\n%w", id, types.ErrBookNotExist)
}

func (i *InMemoryDb) GetUser(id int64) (types.User, error) {
	for _, user := range i.users {
		if user.ID == id {
			return types.CopyUser(user), nil
		}
	}

	return types.User{}, fmt.Errorf("get user with id %d not found;\n%w", id, types.ErrUserNotExist)
}

func (i *InMemoryDb) AddBook(book types.Book) (types.Book, error) {
	fmt.Printf("Start add book %v\n", book)
	bookId++
	book.ID = bookId
	fmt.Printf("Updated id %v\n", book)
	i.books = append(i.books, book)
	return book, nil
}

func (i *InMemoryDb) RemoveBook(id int64) error {
	index := -1

	for i2, book := range i.books {
		if book.ID == id {
			index = i2
			break
		}
	}

	if index == -1 {
		return types.ErrBookNotExist
	}

	i.books = append(i.books[:index], i.books[index+1:]...)

	return nil
}
