package types

import "errors"

var (
	ErrAuthorNotExist = errors.New("author doesn't exist")
)

type Author struct {
	ID        int64  `json:"id"`
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
}
