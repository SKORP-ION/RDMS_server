package main

import (
	"Rostelecom_Device_Management_System/rest"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/getWS", rest.GetWorkstations).Methods("GET")
	r.HandleFunc("/createWS", rest.RegisteringWorkStation).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", r))
}
