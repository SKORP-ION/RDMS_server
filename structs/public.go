package structs

type PublicWorkstation struct {
	Name string `gorm:"colummn:name" json:"name"`
	Serial string `gorm:"column:serial_number"`
	Personal_key string `gorm:"column:personal_key"`
	Token string `json:"token"`
}
func (pws PublicWorkstation) FromWs(ws Workstation) PublicWorkstation {
	pws.Name = ws.Name
	pws.Serial = ws.Serial
	pws.Personal_key = ws.Personal_key
	return pws
}