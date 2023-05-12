package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	models "ani-gin/models"
	"github.com/golodash/galidator"
)


// Formating for JSON validations
var (
	g = galidator.New()
	POSTCustomizer = g.Validator(models.CreateAnimeSchema{})
	PATCHCustomizer = g.Validator(models.UpdateAnimeSchema{})
)


// /api/anime (POST)
func CreateAnime(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var body models.CreateAnimeSchema
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": POSTCustomizer.DecryptErrors(err)})
		return
	}

	new_anime := models.AnimeSchema{Title: body.Title, Country: body.Country, Description: body.Description}
	db.Create(&new_anime)

	c.JSON(http.StatusCreated, gin.H{"data": new_anime})
}


// /api/anime (GET)
func ReadAnime(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var anime []models.AnimeSchema
	db.Find(&anime)

	c.JSON(http.StatusOK, anime)
}


// /api/anime/<int:id> (GET)
func ReadAnimeByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var anime models.AnimeSchema
	if err := db.Where("id = ?", c.Param("id")).First(&anime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ANIME_DOESNT_EXIST"})
		return
	}

	c.JSON(http.StatusOK, anime)
}



// /api/anime/<int:id> (PATCH)
func UpdateAnimeByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var anime models.AnimeSchema
	if err := db.Where("id = ?", c.Param("id")).First(&anime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ANIME_DOESNT_EXIST"})
		return
	}

	var body models.UpdateAnimeSchema
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": PATCHCustomizer.DecryptErrors(err)})
		return
	}

	db.Model(&anime).Updates(body)

	c.JSON(http.StatusOK, gin.H{"data": anime})
}


// /api/anime/<int:id> (DELETE)
func DeleteAnimeByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var anime models.AnimeSchema
	if err := db.Where("id = ?", c.Param("id")).First(&anime).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ANIME_DOESNT_EXIST"})
		return
	}

	db.Delete(&anime)

	c.JSON(http.StatusOK, gin.H{"deleted": true})
}