package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"time"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate UUID
		id := uuid.New().String()
		// Set context variable
		c.Set("x-request-id", id)
		// Set header
		c.Header("x-request-id", id)
		c.Next()
	}
}

// Request logging middleware
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetString("x-request-id")
		ClientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		method := c.Request.Method
		path := c.Request.URL.Path
		t := time.Now()
		c.Next()
		latency := float32(time.Since(t).Seconds())
		status := c.Writer.Status()
		log.Info().Str("Request ID", requestID).Str("Client IP", ClientIP).
			Str("User Agent", userAgent).Str("Method", method).Str("Path", path).
			Float32("Latency", latency).Int("Status", status).Msg("")

	}
}
