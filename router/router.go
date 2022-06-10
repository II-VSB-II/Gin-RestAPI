package router

import (
	v1 "github.com/Gin-RestAPI/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouterV1(Router *gin.RouterGroup) {

	UserRouter := Router.Group("/user")
	{
		UserRouter.POST("/register", v1.Register)
		UserRouter.POST("/login", v1.Login)
	}
}