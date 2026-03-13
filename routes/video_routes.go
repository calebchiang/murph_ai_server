package routes

import (
	"github.com/calebchiang/thirdparty_server/controllers"
	"github.com/calebchiang/thirdparty_server/middleware"
	"github.com/gin-gonic/gin"
)

func VideoRoutes(r *gin.Engine) {

	auth := r.Group("/")
	auth.Use(middleware.RequireAuth())
	{
		auth.POST("/generate_video", controllers.GenerateVideo)
	}

}
