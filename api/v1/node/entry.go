package node

import "github.com/gin-gonic/gin"

type NodeRouter struct {
}

func (nr *NodeRouter) InitNodeRouter(rg *gin.RouterGroup) {
	rg.GET("/list", nr.List)
}

var Entry = new(NodeRouter)
