package main

import (
	"Rostelecom_Device_Management_System/rest"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)


func main() {
	//Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	//Создание роутера и привязывание адресов к функциям
	r := mux.NewRouter()
	r.HandleFunc("/public/authorization", rest.Authorization).Methods("POST")
	r.HandleFunc("/public/workstations/registerWS", rest.RegisteringWorkStation).Methods("POST")
	r.HandleFunc("/private/sysinfo/", rest.AddSysInfo).Methods("POST")
	r.HandleFunc("/private/workstations/getWS", rest.GetWorkstations).Methods("GET")

	addr := os.Getenv("srv_host") + ":" + os.Getenv("srv_port")
	fmt.Printf("Server started at %s", addr)

	//Старт сервиса
	log.Fatal(http.ListenAndServe(addr, r))
}

