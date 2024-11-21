package post

import (
	"context"
	"post-api/internal/model"
	"post-api/pkg"
)

func (s *service) GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error) {
	return s.repository.GetList(ctx, pagination)
}
