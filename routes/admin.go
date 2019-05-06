package routes

import (
	"blog-api-by-gin/controller/admin"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoute(group *gin.RouterGroup) {
	v1 := group.Group("v1")
	{
		v1.GET("/demo", admin.Demo)
	}
}
