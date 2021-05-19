package rest

import (
	"RDMS_server/database"
	"RDMS_server/security"
	"net/http"
)

func GetPackagesList(w http.ResponseWriter, r *http.Request) {
	if status, err := security.JwtAuth(r); !status || err != nil {
		SendResponse(http.StatusUnauthorized, &w, "Unauthorized")
		return
	}

	ws_name := r.Header.Get("Workstation_name")
	packages, err := database.GetPackagesList(ws_name)

	if err != nil {
		SendResponse(http.StatusInternalServerError, &w, "Internal Server Error. Can't send packages list")
		return
	}

	if packages.Len() == 0 {
		SendResponse(http.StatusNoContent, &w, "There are no packages for this workstation")
		return
	}

	SendResponse(http.StatusOK, &w, packages)
}
