package routes

import (
	"blog-api-by-gin/controller/api"
	v1 "blog-api-by-gin/controller/api/v1"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoute(group *gin.RouterGroup) {
	apiv1 := group.Group("v1")
	apiv1.GET("/demo", api.Demo)

	// tags
	apiv1.GET("/tags",v1.Tags)
	apiv1.POST("/tags",v1.AddTag)
	apiv1.PUT("/tags/:id",v1.UpdateTag)
	apiv1.DELETE("/tags/:id",v1.DeleteTag)

	// articles
	apiv1.GET("/articles",v1.Articles)
	apiv1.POST("/articles",v1.AddArticle)
	//apiv1.PUT("/tags/:id",v1.UpdateTag)
	//apiv1.DELETE("/tags/:id",v1.DeleteTag)
}
