package database

import (
	"Rostelecom_Device_Management_System/structs"
	"errors"
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
}

func GetWorkstations() []structs.Workstation {
	var ws_list []structs.Workstation
	//ws := structs.Workstation{}
	db.Raw("SELECT * FROM workstations;").Scan(&ws_list)
	return ws_list
}

func GetWorkstationByName(name string) (structs.Workstation, error) {
	ws := structs.Workstation{}
	db.Raw("SELECT * FROM workstations WHERE name = ?;", name).Scan(&ws)
	if ws.Id == 0 {
		return ws, errors.New("Workstation not found")
	}
	return ws, nil
}

func RegisterWS(ws structs.Workstation) error {
	err := db.Table("workstations").Where("name = ", ws.Name).Update("serial_number", ws.Serial).Error
	return err
}