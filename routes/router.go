package routes

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()

	// register api routes
	api := router.Group("api")
	{
		RegisterApiRoute(api)
	}

	// register admin routes
	admin := router.Group("admin")
	{
		RegisterAdminRoute(admin)
	}
	router.Run()
}