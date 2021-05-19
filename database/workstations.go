package database

import (
	"RDMS_server/structures"
	"RDMS_server/utils"
	"errors"
)

func GetWorkstations() []structures.Workstation {
	var ws_list []structures.Workstation
	//ws := structures.Workstation{}
	//db.Raw("SELECT * FROM workstations;").Scan(&ws_list)
	db.Find(&ws_list)
	return ws_list
}

func GetWorkstationByName(name string) (structures.Workstation, error) {
	ws := structures.Workstation{}
	db.Table("workstations").Where("name = ?", name).First(&ws)
	if ws.Id == 0 {
		return ws, errors.New("Workstation not found")
	}
	return ws, nil
}

func RegisterWS(ws *structures.Workstation) error {
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