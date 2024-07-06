package initialize

import (
	"github.com/gin-gonic/gin"
	"xx/app/middlewares"
	"xx/router"
)

func Routers() *gin.Engine {

	Router := gin.Default()

	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/go")
	router.InitNotifyRouter(ApiGroup)
	router.InitTokenRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	router.InitEsRouter(ApiGroup)

	return Router
}
