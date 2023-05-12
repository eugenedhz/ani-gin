package models


type Genre struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
}

type Type struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Name        string `json"name"`
}

type AnimeSchema struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Country     string `json:"country"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Genre       []Genre // anime might has many genres
	Type        []Type // anime might has many types
}

type CreateAnimeSchema struct {
	Title       string `json:"title" binding:"required,min=2,max=4"`
	Country     string `json:"country" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required,email"`
}

type UpdateAnimeSchema struct {
	Title       string `json:"title"`
	Country     string `json:"country"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type CreateGenreSchema struct {
	Name        string `json"name" binding:"required,min=1"`
}

type UpdateGenreSchema struct {
	Name        string `json"name" binding:"min=1"`
}

type CreateTypeSchema struct {
	Name        string `json"name" binding:"required,min=1"`
}

type UpdateTypeSchema struct {
	Name        string `json"name" binding:"min=1"`
}