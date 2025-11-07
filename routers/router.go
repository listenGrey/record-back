package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"record-back/handlers"
	"time"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	// 配置CORS中间件
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"POST", "GET", "DELETE", "PUT"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	// 使用cors中间件
	r.Use(cors.New(config))

	// 处理空白路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, "error")
	})

	// 根路由
	v1 := r.Group("/api/v1")

	// 测试用
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// 上传合约记录
	v1.POST("/trades", handlers.CreateTrade)

	// 获取统计信息
	// v1.POST("/markets", handlers.GetTrades)

	return r
}
