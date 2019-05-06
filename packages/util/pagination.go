package util

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) int {
	result := 0

	page := com.StrTo(c.DefaultQuery("page", "0")).MustInt()

	if page > 0 {
		result = (page - 1) * GetPageSize(c)
	}
	return result
}

// 获取分页尺寸
func GetPageSize(c *gin.Context) int {
	return com.StrTo(c.DefaultQuery("page_size", "20")).MustInt()
}
