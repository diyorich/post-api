package service

import (
	"context"
	"github.com/diyorich/post-api/internal/model"
	"github.com/diyorich/post-api/pkg"
)

type PostService interface {
	GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error)
	Load(ctx context.Context, filePath string) error
}
