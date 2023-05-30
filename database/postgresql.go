package database

import (
	"assignment-2/entity"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	db *gorm.DB
)

func init() {
	if env := godotenv.Load(); env != nil {
		log.Panic("Error occurred while trying to get env data:", env)
	}

	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error occurred while trying to connect to database:", err)
	}

	if err := db.AutoMigrate(&entity.Order{}, &entity.Item{}); err != nil {
		log.Panic("Error occurred while trying to perform database migrations:", err)
	}
}

func GetDataBaseInstance() *gorm.DB {
	return db
}
