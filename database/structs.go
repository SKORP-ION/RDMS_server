package database

import (
	"RDMS_server/structs"
	"encoding/json"
	"errors"
	"time"
)

type SysinfoModel struct {
	Id uint32 `gorm:"column:ws_id"`
	Info string `gorm:"info"`
	Ts time.Time `gorm:"ts"`
}

func (s SysinfoModel) Create(ws *structs.Workstation, sys *structs.Sysinfo) (SysinfoModel, error) {
	data, err := json.Marshal(sys)

	if err != nil {
		return s, errors.New("Can't parse JSON")
	}

	s.Info = string(data)

	if ws.Id == 0 {
		return s, errors.New("Workstation not found")
	}
	s.Id = ws.Id
	s.Ts = time.Now().UTC()

	return s, nil
}

func (SysinfoModel) TableName() string {
	return "sysinfo"
}

type Config struct {
	Id uint32
	Name string
}

func (Config) TableName() string {
	return "configurations"
}

type Package struct {
	Name       string
	Version    string
	Ord        uint8
	OnServer   bool `gorm:"column:on_server"`
}

func (Package) TableName() string {
	return "packages"
}

type PackagesList struct {
	configPackages []Package //Пакеты, полученные от конфигурации
	rulesPackages []Package // Пакеты, полученные от конкретных правил для пользователя. Имеет приоритет перед конфигурацией
	Packages []Package // Список пакетов для клиента, получается после объединения первых двух
}

//Проверяет наличие пакета с таким же именем, как у аргумента в Packages
func (pl *PackagesList) Has(pkg *Package) (bool, int) {
	for i := 0; i < len(pl.Packages); i++ {
		if pkg.Name == pl.Packages[i].Name {
			return true, i
		}
	}
	return false, 0
}

//Объединяет пакеты с двух списков. В случае конфликта, приоритет имеет пакет из rulesPackages
func (pl *PackagesList) Merge() {
	pl.Packages = pl.configPackages
	for i := 0; i < len(pl.rulesPackages); i++ {
		pkg := pl.rulesPackages[i]
		if status, index := pl.Has(&pkg); status {
			pl.Packages[index] = pkg
		}
	}
}

func (pl *PackagesList) Len() int {
	return len(pl.Packages)
}