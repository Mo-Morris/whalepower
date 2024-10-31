package test

import "github.com/gin-gonic/gin"

type TestRouter struct {
}

func (tr *TestRouter) InitTestRouter(routerGroup *gin.RouterGroup) {
	routerGroup.GET("/ping", tr.Ping)
}

var Entry = new(TestRouter)
