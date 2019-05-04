package routes

import (
	"ginhello/controller/api"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoute(group *gin.RouterGroup) {
	v1 := group.Group("v1")
	{
		v1.GET("/demo", api.Demo)
	}
}
