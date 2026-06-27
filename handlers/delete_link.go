package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteLink(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idstr := c.Param("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "string to convert int"})
		}
		query := "DELETE FROM links WHERE id = $1"

		result, err := db.Exec(query, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete link"})
			return
		}

		rowsaff, _ := result.RowsAffected()
		if rowsaff == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": fmt.Sprintf("link deleted: %d", id)})
	}
}
