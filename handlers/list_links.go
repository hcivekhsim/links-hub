// Package handlers
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hcivekhsim/links-hub/models"
	"github.com/hcivekhsim/links-hub/repository"
)

// GetLinks DI func
func GetLinks(repo repository.LinkRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		links, err := repo.GetAll(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			return
		}
		if links == nil {
			links = []models.Link{}
		}

		c.IndentedJSON(http.StatusOK, links)
	}
}
