package router

import (
	"github.com/gin-gonic/gin"
	"xx/app/api/token"
)

func InitTokenRouter(Route *gin.RouterGroup) {

	notifyRouter := Route.Group("token")
	{
		notifyRouter.POST("", token.GetToken)
	}
}
