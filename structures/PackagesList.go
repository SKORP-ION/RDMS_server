package structures

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
		} else {
			pl.Packages = append(pl.Packages, pkg)
		}
	}
}

func (pl *PackagesList) Len() int {
	return len(pl.Packages)
}

func (pl *PackagesList) SetConfigPackages(packages []Package) {
	pl.configPackages = packages
}

func (pl *PackagesList) SetRulesPackages(packages []Package) {
	pl.rulesPackages = packages
}

func (pl *PackagesList) Sent() []Package {
	return pl.Packages
}