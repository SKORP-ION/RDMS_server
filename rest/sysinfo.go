package rest

import (
	"Rostelecom_Device_Management_System/database"
	"Rostelecom_Device_Management_System/security"
	"Rostelecom_Device_Management_System/structs"
	"encoding/json"
	"net/http"
)

func AddSysInfo(w http.ResponseWriter, r *http.Request) {
	if status, _ := security.JwtAuth(r); !status {
		SendResponse(http.StatusForbidden, &w, "Unauthorized")
		return
	}
	sysinfo := structs.Sysinfo{}
	err := json.NewDecoder(r.Body).Decode(&sysinfo)

	if err != nil {
		SendResponse(http.StatusBadRequest, &w, "Can't parse JSON")
		return
	}

	if sysinfo.Name == "" {
		SendResponse(http.StatusBadRequest, &w, "Missing workstation name")
		return
	}

	ws, err := database.GetWorkstationByName(sysinfo.Name)

	if err != nil {
		SendResponse(http.StatusForbidden, &w, err.Error())
		return
	}

	err = database.PutSysinfo(&ws, &sysinfo)

	if err != nil {
		SendResponse(http.StatusInternalServerError, &w, "An error occerred while trying to load data into the database")
		return
	}

	SendResponse(http.StatusCreated, &w, "Success")
}