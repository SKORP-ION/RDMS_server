package structures

type Package struct {
	Name       string
	Version    string
	Ord        uint8
	OnServer   bool `gorm:"column:on_server"`
}

func (Package) TableName() string {
	return "packages"
}