package router

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ohimma/odemo/gin/config"
	"github.com/ohimma/odemo/gin/controller/user"
	"github.com/ohimma/odemo/gin/handler"
	"github.com/ohimma/odemo/gin/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	gin.SetMode(config.Conf.Mode)

	// 引用公共中间件
	router.Use(
		cors.Default(),
		middleware.RequestLogger(),
	)

	// 404组
	router.NoRoute(func(c *gin.Context) {
		// ctx := common.Context{Ctx: c}
		path := c.Request.URL.Path
		method := c.Request.Method
		middleware.Logger.Warn("404 http request = ", c.Request)
		handler.ResponseFound(c, 404, nil, fmt.Sprintf("%s %s not found", method, path))
	})

	// 测试组
	t := router.Group("/test")
	{
		t.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello ping")
		})
		t.POST("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
	RouterHealth(t)

	// 用户管理 路由组
	router.POST("/api/login", user.UserLogin)

	u := router.Group("/api/user").Use(middleware.Auth())
	// u := router.Group("/api/user")
	{
		u.GET("/user", user.UserList)
		u.PUT("/user", user.UserUpdate)
		u.POST("/user", user.UserCreate)
		u.DELETE("/user", user.UserDelete)

		u.GET("/role", user.RoleList)
		u.PUT("/role", user.RoleUpdate)
		u.POST("/role", user.RoleCreate)
		u.DELETE("/role", user.RoleDelete)

		u.GET("/auth", user.AuthList)
		u.PUT("/auth", user.AuthUpdate)
		u.POST("/auth", user.AuthCreate)
		u.DELETE("/auth", user.AuthDelete)
	}
	return router
}
