package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/setcreed/gin-apiserver/pkg/service"
)

type HelloController interface {
	Hello(c *gin.Context)
}

type helloController struct {
	helloService service.HelloService
}

func NewHelloController() HelloController {
	return &helloController{
		helloService: service.NewHelloService(),
	}
}

func (h *helloController) Hello(c *gin.Context) {
	c.JSON(200, h.helloService.SayHello())
}
