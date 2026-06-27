package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hcivekhsim/links-hub/models"
)

func InsertLink(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		link := models.Link{}

		if err := c.ShouldBindJSON(&link); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		query := `INSERT INTO links (title, url, description) 
		VALUES ($1, $2, $3) RETURNING id `

		var id int
		err := db.QueryRow(query, link.Title, link.URL, link.Desc).Scan(&id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create link in BD"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": fmt.Sprintf("link created: %d", id)})
	}
}
