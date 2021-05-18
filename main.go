package main

import (
	"RDMS_server/rest"
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

	//public URI
	r.HandleFunc("/public/authorization", rest.Authorization).Methods("POST")
	r.HandleFunc("/public/workstations/registerWS", rest.RegisteringWorkStation).Methods("POST")

	//private URI
	r.HandleFunc("/private/sysinfo", rest.AddSysInfo).Methods("POST")

	//URI администратора
	r.HandleFunc("/admin/workstations/getWS", rest.GetWorkstations).Methods("GET")

	addr := os.Getenv("srv_host") + ":" + os.Getenv("srv_port")
	fmt.Printf("Server started at %s", addr)

	//Старт сервиса
	log.Fatal(http.ListenAndServe(addr, r))
	//TODO:Протестить авторизацию по токену
}

