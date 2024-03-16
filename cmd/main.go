package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/setcreed/gin-apiserver/pkg/config"
	"github.com/setcreed/gin-apiserver/pkg/log"
	"github.com/setcreed/gin-apiserver/pkg/routes"
	"github.com/setcreed/gin-apiserver/pkg/utils"
)

func main() {
	utils.SetupSigusr1Trap()
	r := gin.Default()

	m := config.GetString(config.FLAG_KEY_GIN_MODE)
	gin.SetMode(m)

	routes.InstallRoutes(r)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.GetString(config.FLAG_KEY_SERVER_HOST), config.GetInt(config.FLAG_KEY_SERVER_PORT)),
		Handler: r,
	}

	// 优雅退出逻辑
	go func() {
		// 创建一个接收系统信号的通道
		quit := make(chan os.Signal, 1)
		// 监听SIGINT和SIGTERM信号
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		log.Infof("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Server Shutdown: %v", err)
		}
		//log.Infof("Server exiting")
	}()

	log.Infof("Run server at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
