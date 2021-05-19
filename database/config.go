package database

import "RDMS_server/structures"

func GetConfigByWs(ws *structures.Workstation) (structures.Config, error) {
	cfg := structures.Config{}
	err := db.Where("id = ?", ws.Config).First(&cfg).Error
	return cfg, err
}
