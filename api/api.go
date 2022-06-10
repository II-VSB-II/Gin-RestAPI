package api

import (
	middlewares "github.com/II-VSB-II/Gin-RestAPI/middlewares"
	routers "github.com/II-VSB-II/Gin-RestAPI/routers"
	"github.com/II-VSB-II/Gin-RestAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Function to setup the api object
func SetupApp() *gin.Engine {
	log.Info().Msg("Initializing the MLP APIs.")

	// Create barebone engine
	api := gin.New()
	// Add default recovery middleware
	api.Use(gin.Recovery())

	// disabling the trusted proxy feature
	api.SetTrustedProxies(nil)

	// Add cors, request ID and request logging middleware
	log.Info().Msg("Adding CORS, Request ID and Request Logging Middleware.")
	api.Use(middlewares.CORSMiddleware(), middlewares.RequestID(), middlewares.RequestLogger())

	// Setup routers
	log.Info().Msg("Setting up v1 Routers.")
	routers.SetupRoutersV1(api)

	log.Info().Msg("Creating the Database Connection.")
	// Create the database connection
	if dberr := utils.CreateDBConnection(); dberr != nil {
		log.Err(dberr).Msg("Error Occurred While Creating Database Connection.")
	}

	// Auto migrate database
	err := utils.AutoMigrateDB()
	if err != nil {
		log.Err(err).Msg("Error Occurred While Auto Migrating Database.")
	}

	return api
}
