package pkg

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPagination(c *gin.Context) *Pagination {
	var limit, offset, total int
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit == 0 {
		limit = 10
	}

	offset, err = strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
	}

	return &Pagination{
		Limit:  limit,
		Offset: offset,
		Total:  total,
	}
}
