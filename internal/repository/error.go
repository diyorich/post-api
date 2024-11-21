package repository

import "errors"

var (
	ErrSavePost      = errors.New("error on saving post")
	ErrSerialization = errors.New("error on serialization of post")
	ErrDeserialize   = errors.New("error on deserialization of post")
)
