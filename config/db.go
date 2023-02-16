package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/karthiksbh/go-lang-task/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connectdb() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error in loading .env file")
	}

	dbUser := os.Getenv("PGUSER")
	dbPassword := os.Getenv("PGPASSWORD")
	dbName := os.Getenv("PGDATABASE")
	dbHost := os.Getenv("PGHOST")
	dbPort := os.Getenv("PGPORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", dbHost, dbPort, dbUser, dbPassword, dbName)
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
