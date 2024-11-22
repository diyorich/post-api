package service

import (
	"errors"
)

var (
	ErrFileNotFound = errors.New("file with data is not provided")
	ErrInternal     = errors.New("internal error")
	ErrFileClose    = errors.New("error on closing file")
)
