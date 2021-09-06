package rest

import (
	"RDMS_server/database"
	"RDMS_server/security"
	"RDMS_server/structures"
	"encoding/json"
	"net/http"
	. "RDMS_server/logging"
)

func GetPackagesList(w http.ResponseWriter, r *http.Request) {
	if status, err := security.JwtAuth(r); !status || err != nil {
		SendResponse(http.StatusUnauthorized, &w, "Unauthorized")
		Warning.Println("Unauthorized")
		return
	}

	ws_name := r.Header.Get("Workstation_name")
	packages, err := database.GetPackagesList(ws_name)

	if err != nil {
		SendResponse(http.StatusInternalServerError, &w, "Internal Server Error. Can't send packages list")
		Error.Println(err)
		return
	}

	if packages.Len() == 0 {
		SendResponse(http.StatusNoContent, &w, "There are no packages for this workstation")
		Info.Println("No packages for", ws_name)
		return
	}

	SendResponse(http.StatusOK, &w, packages.Sent())
	Info.Println("Packages list sent to", ws_name)
}

func CreateDownloadSession(w http.ResponseWriter, r *http.Request) {
	if status, err := security.JwtAuth(r); !status || err != nil {
		SendResponse(http.StatusUnauthorized, &w, "Unauthorized")
		Warning.Println("Unauthorized")
	}

	pkg := structures.Package{}
	ws_name := r.Header.Get("Workstation_name")

	err := json.NewDecoder(r.Body).Decode(&pkg)

	if err != nil {
		SendResponse(http.StatusBadRequest, &w, "Can't parse JSON")
		Error.Println("Can't parse JSON, wrong format from", ws_name)
		return
	}

	if pkg.Name == "" || pkg.Version == "" {
		SendResponse(http.StatusBadRequest, &w, "Required data is missing")
		Warning.Println("Missing package name or version for", ws_name)
		return
	}

	session, err := database.CreateDownloadSession(pkg)

	if err != err {
		SendResponse(http.StatusInternalServerError, &w, "Internal error")
		Error.Println(err)
		return
	}

	SendResponse(http.StatusCreated, &w, session.ResponseData())
	Info.Println("Download session created for", ws_name)
}