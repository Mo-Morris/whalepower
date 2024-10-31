package test

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func (*TestRouter) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
