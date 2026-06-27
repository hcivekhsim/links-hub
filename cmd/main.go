package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/hcivekhsim/links-hub/handlers"
	"github.com/hcivekhsim/links-hub/pkg"
)

func main() {
	var err error
	db, err := pkg.NewConnDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	log.Println("DB work")
	r := gin.Default()

	r.GET("/links", handlers.GetLinks(db))
	r.POST("/links", handlers.InsertLink(db))
	r.PUT("/links", handlers.UpdateLink(db))
	r.DELETE("/links/:id", handlers.DeleteLink(db))

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %w", err)
	}
}
