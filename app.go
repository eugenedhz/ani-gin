package main

import (
	"ani-gin/controllers"
	"ani-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupModels()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// Grouping endpoints
	v1 := r.Group("/api/v1")
	{
		v1.POST("/anime", controllers.CreateAnime)
		v1.GET("/anime", controllers.ReadAnime)
		v1.GET("/anime/:id", controllers.ReadAnimeByID)
		v1.PATCH("/anime/:id", controllers.UpdateAnimeByID)
		v1.DELETE("/anime/:id", controllers.DeleteAnimeByID)
	}

	r.Run()
}
