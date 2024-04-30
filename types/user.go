package types

import "errors"

var (
	ErrUserNotExist = errors.New("user does not exist")
)

type UserRole string

const (
	VISITOR   UserRole = "VISITOR"
	LIBRARIAN UserRole = "LIBRARIAN"
)

type User struct {
	ID       int64
	Username string
	Role     UserRole
}

func CopyUser(user User) User {
	return User{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}
}
