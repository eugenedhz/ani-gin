package models

import (
	"fmt"
	"ani-gin/constants"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
	prosgret_conname := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", constants.DbHost, constants.DbPort, constants.DbUser, constants.DbName, constants.DbPassword)
	fmt.Println("conname is\t\t", prosgret_conname)

	db, err := gorm.Open("postgres", prosgret_conname)
	if err != nil {
		panic("Connecting to db failed!")
	}

	db.AutoMigrate(&AnimeSchema{})

	return db
}
