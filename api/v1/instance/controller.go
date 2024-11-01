package instance

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查询算力实例列表
func (nr *InstanceRouter) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// 创建算力实例
func (nr *InstanceRouter) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// 停止算力实例
func (nr *InstanceRouter) Stop(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
