package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file", err)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		dbHost, username, dbName, password)

	conn, err := gorm.Open("postgres", dbUri)

	if err != nil {
		log.Fatal("Could not connect to database:\n", err)
	}

	db = conn

	//TODO:Сделать хэширование секретного ключа юзера
	//TODO:Реализвать отправку MD5 хэшсуммы для пакетов с сервера
}