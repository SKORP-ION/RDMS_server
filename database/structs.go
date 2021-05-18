package database

import (
	"Rostelecom_Device_Management_System/structs"
	"encoding/json"
	"errors"
	"time"
)

type SysinfoSql struct {
	Id uint32 `sql:"ws_id"`
	Info string `sql:"info"`
	Ts time.Time `sql:"ts"`
}

func (s SysinfoSql) Create(ws *structs.Workstation, sys *structs.Sysinfo) (error) {
	data, err := json.Marshal(sys)

	if err != nil {
		return errors.New("Can't parse JSON")
	}

	s.Info = string(data)

	if ws.Id == 0 {
		return errors.New("Workstation not found")
	}
	s.Id = ws.Id
	s.Ts = time.Now().UTC()

	return nil
}

type TokenSql struct {
	ws_id uint32
	token string
	ts time.Time
}
