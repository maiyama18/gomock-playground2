package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/maiyama18/gomock-playground/model"
	"github.com/maiyama18/gomock-playground/usecase"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=kyash_user dbname=kyash_local sslmode=disable password= connect_timeout=3")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.City{})
	db.Delete(&model.City{})
	db.Create(&model.City{Name: "tokyo", Population: 2000})
	db.Create(&model.City{Name: "nagoya", Population: 500})
	db.Create(&model.City{Name: "osaka", Population: 1200})

	cityUsecase := usecase.NewCityUsecase(db)
	city, err := cityUsecase.FindByName("tokyo")
	if err != nil {
		panic(err)
	}

	fmt.Println(city)
}
