package main

import (
	"fmt"
	v1 "github.com/II-VSB-II/Gin-RestAPI/api/v1"
	"github.com/II-VSB-II/Gin-RestAPI/config"
	"github.com/II-VSB-II/Gin-RestAPI/initialize"
	"log"
)

func main() {

	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("cannot load config:", err)
	}

	/*
		docs.SwaggerInfo.Title = "Machine Learning Platform"
		docs.SwaggerInfo.Description = "Machine Learning Platform for Training, Tracking, Deploying and Monitoring AI/ML models."
		docs.SwaggerInfo.Version = "1.0"
		docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}

		err = initialize.InitMySQL()
		if err != nil {
			panic(err)
		}
		defer initialize.Close()
	*/

	r := initialize.SetupRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/version", v1.Version)
	r.Run(fmt.Sprintf(":%d", config.Application.HTTPPort))
}
