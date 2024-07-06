package router

import (
	"github.com/gin-gonic/gin"
	"xx/app/api/es"
)

func InitEsRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("es")
	{
		UserRouter.GET("", es.Add)
	}
}
