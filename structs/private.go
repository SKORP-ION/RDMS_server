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

type Sysinfo struct {
	Name string `json:"name"`
	Model string `json:"model"`
	Cpu string `json:"cpu"`
	Ram string `json:"ram"`
	MBSerial string `json:"MBSerial"`
	Arch string `json:"arch"`
	Core string `json:"core"`
	Os string `json:"os"`
	RDMS_Version string `json:"RMDS_version"`
	Anydesk string `json:"anydesk"`
}
