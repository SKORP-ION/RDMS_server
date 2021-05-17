package rest

import (
	"Rostelecom_Device_Management_System/database"
	"net/http"
)

func GetWorkstations(w http.ResponseWriter, r *http.Request) {
	result := database.GetWorkstations()
	//w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(result)
	sendResponse(http.StatusOK, &w, result)
}