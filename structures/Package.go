package structures

type Package struct {
	Name       string
	Version    string
	Ord        uint8
	OnServer   bool `gorm:"column:on_server"`
	Md5	       string
	Action     uint8
}

func (Package) TableName() string {
	return "packages"
}