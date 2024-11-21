package post

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"post-api/internal/model"
	repoErr "post-api/internal/repository"
	"post-api/internal/repository/cache"
	"post-api/pkg"
)

const sortedPosts = "sorted_post_set"

type repository struct {
	cache *cache.Cache
}

func NewRepository(cache *cache.Cache) *repository {
	return &repository{cache: cache}
}

func (r *repository) Add(ctx context.Context, post model.Post) error {
	const op = "postRepository.Add"

	data, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("%s: %w", op, repoErr.ErrSerialization)
	}

	err = r.cache.ZAdd(ctx, sortedPosts, redis.Z{Score: float64(post.ID), Member: data}).Err()
	if err != nil {
		return repoErr.ErrSavePost
	}

	return nil
}

func (r *repository) GetList(ctx context.Context, pagination *pkg.Pagination) ([]model.Post, error) {
	const op = "postRepository.GetList"
	start := pagination.Offset
	stop := pagination.Offset + pagination.Limit - 1

	data, err := r.cache.ZRangeWithScores(ctx, sortedPosts, int64(start), int64(stop)).Result()
	if err != nil {
		return nil, err
	}

	converted := make([]model.Post, len(data))
	for index, z := range data {
		var post model.Post
		err := json.Unmarshal([]byte(z.Member.(string)), &post)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, repoErr.ErrDeserialize)
		}

		converted[index] = post
	}

	return converted, nil
}
