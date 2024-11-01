package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查询算力节点
func (nr *NodeRouter) List(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
