package rest

import (
	"Rostelecom_Device_Management_System/database"
	"Rostelecom_Device_Management_System/structs"
	"encoding/json"
	"net/http"
)

func RegisteringWorkStation(w http.ResponseWriter, r *http.Request) {
	var recv structs.Workstation

	err := json.NewDecoder(r.Body).Decode(&recv)
	if err != nil {
		sendResponse(http.StatusBadRequest, &w, "Can't parse")
		return
	}

	ws, err := database.GetWorkstationByName(recv.Name)
	if err != nil && err.Error() == "Workstation not found" {
		sendResponse(http.StatusNoContent, &w, "Workstation not found")
		return
	}

	if ws.Serial != "" {
		sendResponse(http.StatusForbidden, &w, "Workstation with this serial number is already exists")
		return
	}

	err = database.RegisterWS(ws)
	if err != nil {
		sendResponse(http.StatusInternalServerError, &w, "Internal error")
		return
	}

	sendResponse(http.StatusOK, &w, ws)
}