package database

import (
	"RDMS_server/structs"
	"encoding/json"
	"errors"
	"time"
)

type SysinfoModel struct {
	Id uint32 `gorm:"column:ws_id"`
	Info string `gorm:"info"`
	Ts time.Time `gorm:"ts"`
}

func (s SysinfoModel) Create(ws *structs.Workstation, sys *structs.Sysinfo) (SysinfoModel, error) {
	data, err := json.Marshal(sys)

	if err != nil {
		return s, errors.New("Can't parse JSON")
	}

	s.Info = string(data)

	if ws.Id == 0 {
		return s, errors.New("Workstation not found")
	}
	s.Id = ws.Id
	s.Ts = time.Now().UTC()

	return s, nil
}

func (SysinfoModel) TableName() string {
	return "sysinfo"
}

type TokenSql struct {
	ws_id uint32
	token string
	ts time.Time
}
