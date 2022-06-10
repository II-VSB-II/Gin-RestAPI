package v1

import (
	"net/http"
)

// Vesrion doc
// @Description  Version
// @Tags  Version
// @Summary Get the version of the platform.
// @Produce  json
// @Success   200      {string}  string  "{"status": "success", "message": null , "data": { "version": "1.0" }}"
// @Router    /version [get]
func Version(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": null,
		"data": {
			"version": "1.0",
		},
	})
}
