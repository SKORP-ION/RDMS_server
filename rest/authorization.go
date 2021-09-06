package rest

import (
	"RDMS_server/database"
	"RDMS_server/security"
	"RDMS_server/structures"
	"encoding/json"
	"net/http"
	. "RDMS_server/logging"
)

func Authorization(w http.ResponseWriter, r *http.Request) {
	var recv structures.Workstation

	err := json.NewDecoder(r.Body).Decode(&recv)
	if err != nil {
		SendResponse(http.StatusBadRequest, &w, "Can't parse json")
		Warning.Println("Can't parse json. Wrong format.")
		return
	}

	ws, err := database.GetWorkstationByName(recv.Name) //Выгрузка информации о машине по имени
	if err != nil && err.Error() == "Workstation not found" {
		SendResponse(http.StatusNoContent, &w, "Workstation not found")
		Info.Println("Workstation", ws.Name, "not found on server")
		return
	}
	if ws.Personal_key == recv.Personal_key {
		token, err := security.CreateToken(ws) //Генерация нового токена для клиента
		if err != nil {
			SendResponse(http.StatusInternalServerError, &w, "Internal Error")
			Error.Println(err)
			return
		}
		SendResponse(http.StatusOK, &w, token) //Отправить токен клиенту
		Info.Println("Token sent to", ws.Name)
	}
}