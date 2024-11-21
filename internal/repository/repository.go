package repository

import (
	"context"
	"post-api/internal/model"
	"post-api/pkg"
)

type PostRepository interface {
	GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error)
	Add(ctx context.Context, post model.Post) error
}
