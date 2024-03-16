/*
Copyright 2023 The gin-apiserver Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
