package main

import (
	"github.com/gin-gonic/gin"
)
import "net/http"
import "ginhello/routes"

func demo(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"success",
		"data":make(map[string]string),
	})
}
func main() {
	routes.Init()
	//routes.AdminInit()
}
