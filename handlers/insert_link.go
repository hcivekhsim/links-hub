package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcivekhsim/links-hub/models"
	"github.com/hcivekhsim/links-hub/repository"
)

func InsertLink(repo repository.LinkRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		link := models.Link{}

		if err := c.ShouldBindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := repo.AddLink(c.Request.Context(), link)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create link in BD"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": fmt.Sprintf("link created: %d", id)})
	}
}
