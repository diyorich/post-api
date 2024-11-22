package handler

import (
	"github.com/diyorich/post-api/pkg"
	"github.com/gin-gonic/gin"
)

type Meta struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func ResponseOK(data interface{}, pagination *pkg.Pagination) gin.H {
	return gin.H{
		"data": data,
		"meta": Meta{
			Limit:  pagination.Limit,
			Offset: pagination.Offset,
			Total:  pagination.Total,
		},
	}
}

func ResponseErr(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
