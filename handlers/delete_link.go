package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hcivekhsim/links-hub/repository"
)

func DeleteLink(repo repository.LinkRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "string to convert int"})
		}

		err = repo.RemoveLink(c.Request.Context(), id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete link"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("link deleted: %d", id)})
	}
}
