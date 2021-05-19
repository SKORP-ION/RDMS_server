package database

import "RDMS_server/structures"

func PutSysinfo(ws *structures.Workstation, info *structures.Sysinfo) error {
	s := structures.SysinfoModel{}
	err := s.Create(ws, info)

	if err != nil {
		return err
	}

	err = db.Create(&s).Error

	if err != nil {
		return err
	}

	return nil
}