package health

import (
	"github.com/gin-gonic/gin"
	"github.com/ohimma/odemo/gin/handler"
)

func HealthCheck(c *gin.Context) {
	handler.ResponseOK(c, 200, "health", "success")
	// c.String(200, "health")
}

func DiskCheck(c *gin.Context) {
	data := "disk"
	handler.ResponseOK(c, 200, data, "success")
}

func CpuCheck(c *gin.Context) {
	data := "cpu"
	handler.ResponseOK(c, 200, data, "success")
}

func MemCheck(c *gin.Context) {
	data := "mem"
	handler.ResponseOK(c, 200, data, "success")
}
