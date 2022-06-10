package routers

import (
	controllers "github.com/II-VSB-II/Gin-RestAPI/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutersV1(app *gin.Engine) {
	v1 := app.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Ping)
		v1.GET("/version", controllers.Version)
		v1.POST("/register", controllers.RegisterUser)
		//v1.GET("/user/:id", controllers.GetUser)
		//v1.PATCH("/user/:id", controllers.UpdateUser)
		//v1.DELETE("/user/:id", controllers.DeleteUser)
	}
}
