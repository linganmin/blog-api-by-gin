package routes

import (
	"ginhello/controller/admin"
	"github.com/gin-gonic/gin"
)

func RegisterAdminRoute(group *gin.RouterGroup) {
	v1 := group.Group("v1")
	{
		v1.GET("/demo", admin.Demo)
	}
}
