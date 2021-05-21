package rest

import (
	"RDMS_server/database"
	"RDMS_server/security"
	"RDMS_server/structures"
	"encoding/json"
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

func CreateDownloadSession(w http.ResponseWriter, r *http.Request) {
	if status, err := security.JwtAuth(r); !status || err != nil {
		SendResponse(http.StatusUnauthorized, &w, "Unauthorized")
	}

	pkg := structures.Package{}

	err := json.NewDecoder(r.Body).Decode(&pkg)

	if err != nil {
		SendResponse(http.StatusBadRequest, &w, "Can't parse JSON")
		return
	}

	if pkg.Name == "" || pkg.Version == "" {
		SendResponse(http.StatusBadRequest, &w, "Required data is missing")
		return
	}

	session, err := database.CreateDownloadSession(pkg)

	if err != err {
		SendResponse(http.StatusInternalServerError, &w, "Internal error")
		return
	}

	SendResponse(http.StatusCreated, &w, session.ResponseData())
}