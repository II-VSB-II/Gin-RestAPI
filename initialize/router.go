package initialize

import (
	"github.com/II-VSB-II/Gin-RestAPI/middleware"
	"github.com/II-VSB-II/Gin-RestAPI/router"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(middleware.LoggerToFile())
	Group := r.Group("api/v1")
	{
		router.InitRouterV1(Group)
	}
	return r
}
