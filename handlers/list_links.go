// Package handlers
package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hcivekhsim/links-hub/models"
)

// GetLinks DI func
func GetLinks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		links := []models.Link{}

		rows, err := db.QueryContext(c.Request.Context(), "SELECT id, title, url, description, created_at, updated_at FROM  links")
		if err != nil {
			log.Fatal("err_select")
		}

		defer rows.Close()

		for rows.Next() {
			l := models.Link{}

			err := rows.Scan(&l.ID, &l.Title, &l.URL, &l.Desc, &l.CreatedAt, &l.UpdatedAt)
			if err != nil {
				log.Fatal("err_scan")
			}
			links = append(links, l)
		}

		c.IndentedJSON(http.StatusOK, links)
	}
}
