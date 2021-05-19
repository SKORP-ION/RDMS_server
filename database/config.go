package database

import "RDMS_server/structs"

func GetConfigByWs(ws *structs.Workstation) (Config, error) {
	cfg := Config{}
	err := db.Where("id = ?", ws.Config).First(&cfg).Error
	return cfg, err
}
