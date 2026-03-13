package controllers

import (
	"net/http"

	"github.com/calebchiang/thirdparty_server/services"
	"github.com/gin-gonic/gin"
)

func GenerateVideo(c *gin.Context) {

	var input struct {
		ImageURL string `json:"image_url"`
		Prompt   string `json:"prompt"`
		Duration int    `json:"duration"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	if input.ImageURL == "" || input.Prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "image_url and prompt are required",
		})
		return
	}

	// default duration if not provided
	if input.Duration == 0 {
		input.Duration = 5
	}

	videoID, err := services.GenerateRunwayVideo(
		input.ImageURL,
		input.Prompt,
		input.Duration,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"video_id": videoID,
	})
}
