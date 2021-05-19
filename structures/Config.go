package structures

type Config struct {
	Id uint32
	Name string
}

func (Config) TableName() string {
	return "configurations"
}
