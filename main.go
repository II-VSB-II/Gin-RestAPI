package main

import (
	"fmt"

	api "github.com/II-VSB-II/Gin-RestAPI/api"
	utils "github.com/II-VSB-II/Gin-RestAPI/utils"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func init() {
	// Set gin mode
	mode := utils.GetEnvVar("API_MODE")
	gin.SetMode(mode)
}

func main() {
	// Setup the api
	api := api.SetupApp()

	// Read ADDR and port
	address := utils.GetEnvVar("API_ADDRESS")
	port := utils.GetEnvVar("API_PORT")
	https := utils.GetEnvVar("API_HTTPS")

	// HTTPS mode
	if https == "true" {
		certFile := utils.GetEnvVar("SSL_CERT")
		certKey := utils.GetEnvVar("SSL_KEY")
		log.Info().Msgf("Starting Machine Learning Platform API on https//:%s:%s", address, port)

		if err := api.RunTLS(fmt.Sprintf("%s:%s", address, port), certFile, certKey); err != nil {
			log.Fatal().Err(err).Msg("Error when starting the Machine Learning Platform API on HTTPS.")
		}
	}
	// HTTP mode
	log.Info().Msgf("Starting Machine Learning Platform API on https//:%s:%s", address, port)
	if err := api.Run(fmt.Sprintf("%s:%s", address, port)); err != nil {
		log.Fatal().Err(err).Msg("Error when starting the Machine Learning Platform API on HTTP.")
	}

}
