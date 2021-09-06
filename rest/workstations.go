package rest

import (
	"RDMS_server/database"
	"RDMS_server/structures"
	"encoding/json"
	"net/http"
)

func RegisteringWorkStation(w http.ResponseWriter, r *http.Request) {
	var recv structures.Workstation

	err := json.NewDecoder(r.Body).Decode(&recv)
	if err != nil {
		SendResponse(http.StatusBadRequest, &w, "Can't parse")
		return
	}

	ws, err := database.GetWorkstationByName(recv.Name)
	if err != nil && err.Error() == "Workstation not found" {
		SendResponse(http.StatusNoContent, &w, "Workstation not found")
		return
	}

	if ws.Serial != "" {
		SendResponse(http.StatusForbidden, &w, "Workstation with this serial number is already exists")
		return
	}

	if status, err := database.IsWSnotRegistered(&ws); !status {
		SendResponse(http.StatusForbidden, &w, "Workstation with this name already exists")
		return
	} else if err != nil {
		SendResponse(http.StatusInternalServerError, &w, "Internal error")
		return
	}

	err = database.RegisterWS(&recv)
	if err != nil {
		SendResponse(http.StatusInternalServerError, &w, "Internal error")
		return
	}

	publicWS := structures.PublicWorkstation{}.FromWs(recv)

	SendResponse(http.StatusOK, &w, publicWS)
}