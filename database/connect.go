package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"ramusa/models"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=postgres password=3242 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err.Error())
		panic("Veritabanına bağlanmadın looo!")
	}
	DB = database
	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
}
