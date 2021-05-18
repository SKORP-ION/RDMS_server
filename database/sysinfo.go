package database

import "Rostelecom_Device_Management_System/structs"

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