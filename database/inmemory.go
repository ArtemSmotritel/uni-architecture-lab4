package database

import (
	"fmt"
	"github.com/artemsmotritel/uni-architecture-lab4/types"
	"slices"
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
		if len(filters.Statuses) == 0 || slices.Contains[[]types.BookStatus](filters.Statuses, book.Status) {
			res = append(res, book)
		}
	}

	return res, nil
}

func (i *InMemoryDb) GetBook(id int64) (types.Book, error) {
	for _, book := range i.books {
		if book.ID == id {
			return book, nil
		}
	}

	return types.Book{}, fmt.Errorf("get book with id %d not found;\n%w", id, types.ErrBookNotExist)
}

func (i *InMemoryDb) GetUser(id int64) (types.User, error) {
	for _, user := range i.users {
		if user.ID == id {
			return user, nil
		}
	}

	return types.User{}, fmt.Errorf("get user with id %d not found;\n%w", id, types.ErrUserNotExist)
}

func (i *InMemoryDb) AddBook(book types.Book) (types.Book, error) {
	bookId++
	book.ID = bookId
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

func (i *InMemoryDb) LendBook(id int64) error {
	for i2 := range i.books {
		if i.books[i2].ID == id {
			if i.books[i2].Status == types.AVAILABLE {
				i.books[i2].Status = types.LENT
				return nil
			} else {
				return types.ErrBookStatusConflict
			}
		}
	}

	return types.ErrBookNotExist
}

func (i *InMemoryDb) ReturnBook(id int64) error {
	for i2 := range i.books {
		if i.books[i2].ID == id {
			if i.books[i2].Status == types.LENT {
				i.books[i2].Status = types.AVAILABLE
				return nil
			} else {
				return types.ErrBookStatusConflict
			}
		}
	}

	return types.ErrBookNotExist
}

func (i *InMemoryDb) AddAuthor(author types.Author) error {
	authorId++
	i.authors = append(i.authors, types.Author{
		ID:        authorId,
		FullName:  author.FullName,
		ShortName: author.ShortName,
	})

	return nil
}

func (i *InMemoryDb) GetAuthors() []types.Author {
	return i.authors
}

func (i *InMemoryDb) GetAuthor(id int64) (types.Author, error) {
	for _, author := range i.authors {
		if author.ID == id {
			return author, nil
		}
	}

	return types.Author{}, types.ErrAuthorNotExist
}
