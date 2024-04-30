package service

import (
	"errors"
	"github.com/artemsmotritel/uni-architecture-lab4/database"
)

var (
	ErrFailedValidation = errors.New("failed validation")
)

type Service struct {
	DB database.Database
}

func NewService(db database.Database) *Service {
	return &Service{
		DB: db,
	}
}
