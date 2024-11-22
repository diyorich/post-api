package handler

import (
	"github.com/diyorich/post-api/pkg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *handler) GetList(c *gin.Context) {
	p := pkg.GetPagination(c)

	data, err := h.postService.GetList(c.Request.Context(), p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseErr(err))
		return
	}

	c.JSON(http.StatusOK, ResponseOK(data, p))
}
