package service

import (
	"context"
	"post-api/internal/model"
	"post-api/pkg"
)

type PostService interface {
	GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error)
	Load(ctx context.Context, filePath string) error
}
