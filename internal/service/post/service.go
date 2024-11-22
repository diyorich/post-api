package post

import "github.com/diyorich/post-api/internal/repository"

type service struct {
	repository repository.PostRepository
}

func NewService(repository repository.PostRepository) *service {
	return &service{repository: repository}
}
