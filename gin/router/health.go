package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ohimma/odemo/gin/controller/health"
)

func RouterHealth(group *gin.RouterGroup) {
	router := group.Group("")
	{
		router.GET("/health", health.HealthCheck)
		router.GET("/disk", health.DiskCheck)
		router.GET("/cpu", health.CpuCheck)
		router.GET("/mem", health.MemCheck)
	}
}
