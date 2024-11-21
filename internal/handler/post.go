package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"post-api/internal/service"
	"post-api/pkg"
)

type handler struct {
	postService service.PostService
}

func NewHandler(service service.PostService) *handler {
	return &handler{postService: service}
}

func (h *handler) GetList(c *gin.Context) {
	p := pkg.GetPagination(c)

	data, err := h.postService.GetList(c.Request.Context(), p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
