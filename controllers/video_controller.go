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
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	videoURL, err := services.GenerateRunwayVideo(input.ImageURL, input.Prompt)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"video_url": videoURL,
	})
}
