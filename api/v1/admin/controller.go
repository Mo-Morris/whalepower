package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册算力节点
func (ar *AdminRouter) RegistryNode(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
