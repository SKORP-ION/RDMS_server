package database

import "RDMS_server/structs"

func PutSysinfo(ws *structs.Workstation, info *structs.Sysinfo) error {
	s := SysinfoModel{}
	s, err := s.Create(ws, info)

	if err != nil {
		return err
	}

	err = db.Create(&s).Error

	if err != nil {
		return err
	}

	return nil
}