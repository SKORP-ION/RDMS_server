package rest

import (
	"Rostelecom_Device_Management_System/database"
	"Rostelecom_Device_Management_System/security"
	"Rostelecom_Device_Management_System/structs"
	"encoding/json"
	"net/http"
)

func Authorization(w http.ResponseWriter, r *http.Request) {
	var recv structs.Workstation

	err := json.NewDecoder(r.Body).Decode(&recv)
	if err != nil {
		SendResponse(http.StatusBadRequest, &w, "Can't parse json")
		return
	}

	ws, err := database.GetWorkstationByName(recv.Name) //Выгрузка информации о машине по имени
	if err != nil && err.Error() == "Workstation not found" {
		SendResponse(http.StatusNoContent, &w, "Workstation not found")
		return
	}
	if ws.Personal_key == recv.Personal_key {
		token, err := security.CreateToken(ws) //Создание записи в БД о токене
		if err != nil {
			SendResponse(http.StatusInternalServerError, &w, "Internal Error")
			return
		}
		SendResponse(http.StatusOK, &w, token) //Отправить токен клиенту
	}
}