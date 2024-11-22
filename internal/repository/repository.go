package repository

import (
	"context"
	"github.com/diyorich/post-api/internal/model"
	"github.com/diyorich/post-api/pkg"
)

type PostRepository interface {
	GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error)
	Add(ctx context.Context, post model.Post) error
}
