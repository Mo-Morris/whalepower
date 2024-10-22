package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestRouter struct {
	ApiGroup *gin.RouterGroup
}

func (*TestRouter) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
