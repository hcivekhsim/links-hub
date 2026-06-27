package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hcivekhsim/links-hub/models"
)

func UpdateLink(db *sql.DB) gin.HandlerFunc {
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
		query := `UPDATE links SET 
		title = COALESCE(NULLIF($1,''), title), url = COALESCE(NULLIF($2,''), url), description = COALESCE(NULLIF($3,''), description), 
		updated_at = NOW() WHERE id = $4; `

		result, err := db.Exec(query, link.Title, link.URL, link.Desc, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update link"})
			return
		}

		rowsaff, _ := result.RowsAffected()
		if rowsaff == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"status": fmt.Sprintf("link updated: %d", id)})
	}
}
