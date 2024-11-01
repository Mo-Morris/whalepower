package instance

import "github.com/gin-gonic/gin"

type InstanceRouter struct {
}

func (nr *InstanceRouter) InitNodeRouter(rg *gin.RouterGroup) {
	rg.GET("/list", nr.List)
}

var Entry = new(InstanceRouter)
