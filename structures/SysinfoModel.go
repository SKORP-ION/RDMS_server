package structures

import (
	"encoding/json"
	"errors"
	"time"
)

type SysinfoModel struct {
	Id uint32 `gorm:"column:ws_id"`
	//Info string `gorm:"info"`
	Info string `gorm:"column:info"`
	Ts time.Time `gorm:"ts"`
}

func (s *SysinfoModel) Create(ws *Workstation, sys *Sysinfo) error {
	data, _ := json.Marshal(sys)

	s.Info = string(data)

	if ws.Id == 0 {
		return errors.New("Workstation not found")
	}

	s.Id = ws.Id
	s.Ts = time.Now().UTC()
	return nil
}

func (SysinfoModel) TableName() string {
	return "sysinfo"
}
