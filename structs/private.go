package structs

type Workstation struct {
	Id uint32 `gorm:"column:id"`
	Name string `gorm:"colummn:name" json:"name"`
	Serial string `gorm:"column:serial_number"`
	Group uint16 `gorm:"column:group"`
	Config uint16 `gorm:"column:config"`
	Separate_rules uint32 `gorm:"column:separate_rules"`
	Personal_key string `gorm:"column:personal_key"`
	Description string `gorm:"column:description"`
}