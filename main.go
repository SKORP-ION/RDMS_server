package main

import (
	"Rostelecom_Device_Management_System/app"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.JwtAuthentication)

	port := os.Getenv("srv_port")

	fmt.Printf("Server started at localhost:%s", port)

	err := http.ListenAndServe(":" + port, router)

	if err != nil {
		fmt.Print(err)
	}

}
