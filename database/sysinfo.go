package database

import "Rostelecom_Device_Management_System/structs"

func PutSysinfo(ws *structs.Workstation, info *structs.Sysinfo) error {
	s := SysinfoSql{}
	err := s.Create(ws, info)

	if err != nil {
		return err
	}

	err = db.Table("sysinfo").Create(s).Error

	if err != nil {
		return err
	}

	return nil
}