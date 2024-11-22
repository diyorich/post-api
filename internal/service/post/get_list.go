package post

import (
	"context"
	"github.com/diyorich/post-api/internal/model"
	"github.com/diyorich/post-api/pkg"
)

func (s *service) GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error) {
	return s.repository.GetList(ctx, pagination)
}
