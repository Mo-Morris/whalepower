package admin

import "github.com/gin-gonic/gin"

type AdminRouter struct {
}

func (ar *AdminRouter) InitAdminRouter(rg *gin.RouterGroup) {
	rg.POST("/node/regist", ar.RegistryNode)

}

var Entry = new(AdminRouter)
