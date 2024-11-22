package handler

import "github.com/diyorich/post-api/internal/service"

type handler struct {
	postService service.PostService
}

func NewHandler(service service.PostService) *handler {
	return &handler{postService: service}
}
