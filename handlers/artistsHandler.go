package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"

	models "go_word_cloud/models"
)

func ArtistsIndex(c *gin.Context) {
	var artists = []models.Artist{
		{ID: 1, Name: "The Beatles"},
		{ID: 2, Name: "The Rolling Stones"},
		{ID: 3, Name: "The Who"},
		{ID: 4, Name: "Led Zeppelin"},
		{ID: 5, Name: "Pink Floyd"},
		{ID: 6, Name: "Bob Dylan"},
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Artists Index",
		"data":    artists,
	})
}