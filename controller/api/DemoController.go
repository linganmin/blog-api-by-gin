package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Demo(c *gin.Context)  {
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"前台接口",
		"data":make(map[string]string),
	})
}