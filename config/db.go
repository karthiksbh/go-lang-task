package config

import (
	"fmt"
	"log"

	"github.com/karthiksbh/go-lang-task/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connectdb() {
	dsn := "host=localhost user=postgres password=postgres dbname=go-langtask port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connection established")
	db.AutoMigrate(&models.ProceedingAdd{})
	db.AutoMigrate(&models.Property_vals{})
	db.AutoMigrate(&models.AddressInformation{})
	db.AutoMigrate(&models.Property_Information{})

	db.AutoMigrate(&models.Party_vals{})
	db.AutoMigrate(&models.Prosecution_Entry{})
	db.AutoMigrate(&models.Proceeding_Entry{})

	DB = db
	fmt.Println("Database migration completed")
}
