package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	env := godotenv.Load()
	if env != nil {
		log.Fatalln("env not found")
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbType := os.Getenv("db_type")

	dbUri := fmt.Sprintf("host=%s user%s dbname=%s sslmode=disable password=%s",
		dbHost, username, dbName, password)

	conn, err := gorm.Open(dbType, dbUri)

	if err != nil {
		log.Fatal(err)
	}

	db = conn

}

func GetDB() *gorm.DB {
	return db
}