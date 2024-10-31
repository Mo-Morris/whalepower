package node

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 查询算力节点
func (nr *NodeRouter) NodeList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
