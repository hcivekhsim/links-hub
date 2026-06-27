package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hcivekhsim/links-hub/models"
	"github.com/hcivekhsim/links-hub/repository"
)

func UpdateLink(repo repository.LinkRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		link := models.LinkUpdate{}

		if err := c.ShouldBindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		idstr := c.Query("id")

		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to convert int"})
			return
		}

		err = repo.ChangeLink(c.Request.Context(), id, link)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update link"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": fmt.Sprintf("link updated: %d", id)})
	}
}
