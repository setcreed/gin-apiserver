package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/setcreed/gin-apiserver/pkg/controller"
	"github.com/setcreed/gin-apiserver/pkg/middleware"
)

func InstallRoutes(r *gin.Engine) {
	r.Use(gin.Recovery())

	r.Use(middleware.Cors())

	r.GET("/version", controller.Version)

	rootGroup := r.Group("/api/v1")
	rootGroup.Use(middleware.BasicAuthMiddleware())

	{
		rootGroup.GET("/ping", controller.Ping)

		hello := controller.NewHelloController()
		rootGroup.GET("/hello", hello.Hello)
	}
}
