package database

import (
	"RDMS_server/structs"
	"RDMS_server/utils"
	"errors"
)

func GetWorkstations() []structs.Workstation {
	var ws_list []structs.Workstation
	//ws := structs.Workstation{}
	db.Raw("SELECT * FROM workstations;").Scan(&ws_list)
	return ws_list
}

func GetWorkstationByName(name string) (structs.Workstation, error) {
	ws := structs.Workstation{}
	db.Raw("SELECT * FROM workstations WHERE name = ?;", name).Scan(&ws)
	if ws.Id == 0 {
		return ws, errors.New("Workstation not found")
	}
	return ws, nil
}

func RegisterWS(ws *structs.Workstation) error {
	ws.Personal_key = utils.GeneratePassword()
	err := db.Table("workstations").Where("name = ?", ws.Name).Updates(ws).Error
	return err
}

func WorkstationAvailability(name string) bool {
	ws, err:= GetWorkstationByName(name)

	if err != nil {
		return false
	}

	if ws.Id == 0 {
		return false
	}

	return true
}