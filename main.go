package main

import (
	. "RDMS_server/logging"
	"RDMS_server/rest"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)


func main() {
	//Загрузка переменных окружения
	err := godotenv.Load()
	if err != nil {
		Error.Fatal(err)
	}

	//Создание роутера и привязывание адресов к функциям
	r := mux.NewRouter()


	//public URI
	r.HandleFunc("/public/authorization", rest.Authorization).Methods("POST")
	r.HandleFunc("/public/workstations/registerWS", rest.RegisteringWorkStation).Methods("POST")

	//private URI
	r.HandleFunc("/private/sysinfo/putInfo", rest.AddSysInfo).Methods("POST")
	r.HandleFunc("/private/packages/getPackagesList", rest.GetPackagesList).Methods("GET")
	r.HandleFunc("/private/packages/getSessionKey", rest.CreateDownloadSession).Methods("POST")

	//URI администратора
	//r.HandleFunc("/admin/workstations/getWS", rest.GetWorkstations).Methods("GET")



	addr := os.Getenv("srv_host") + ":" + os.Getenv("srv_port")
	Info.Printf("Server started at %s\n", addr)

	//Старт сервиса
	Error.Fatal(http.ListenAndServe(addr, r))

	//TODO:Реализовать создание сессии на скачивание по запросу
}

