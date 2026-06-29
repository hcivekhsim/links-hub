package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hcivekhsim/links-hub/handlers"
	"github.com/hcivekhsim/links-hub/pkg"
	"github.com/hcivekhsim/links-hub/repository"
)

func main() {
	var err error
	db, err := pkg.NewConnDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	linkRepo := repository.NewPostgresRepo(db)

	log.Println("DB work")

	r := gin.Default()

	r.LoadHTMLFiles("components/index.html")

	r.Static("/css", "./components/css")
	r.Static("/font", "./components/font")
	r.Static("/js", "./components/js")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/links/", handlers.GetLinks(linkRepo))
	r.POST("/links/", handlers.InsertLink(linkRepo))
	r.PUT("/links/", handlers.UpdateLink(linkRepo))
	r.DELETE("/links/:id", handlers.DeleteLink(linkRepo))

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %w", err)
	}
}
