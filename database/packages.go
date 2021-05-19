package database

import "RDMS_server/structures"

func GetPackagesList(name string) (structures.PackagesList, error) {
	packages := structures.PackagesList{}
	ws, err := GetWorkstationByName(name)

	if err != nil {
		return packages, err
	}

	config, err := GetConfigByWs(&ws)

	if err != nil {
		return packages, err
	}

	config_packages, err := getConfigPackagesList(config.Id)

	if err != nil {
		return packages, err
	}

	rules_packages, err := getRulesPackagesList(ws.Separate_rules)

	if err != nil {
		return packages, err
	}

	packages.SetConfigPackages(config_packages)
	packages.SetRulesPackages(rules_packages)
	packages.Merge()

	return packages, nil
}

func getConfigPackagesList(cfg_id uint32) ([]structures.Package, error) {
	var packages []structures.Package
	err := db.Joins("INNER JOIN configuration_packages as cpg ON packages.id = cpg.package_id").
		Select("packages.name, packages.version, packages.ord, packages.on_server").
		Where("cpg.config_id = ?", cfg_id).
		Find(&packages).Error
	return packages, err
}

func getRulesPackagesList(rules_id uint32) ([]structures.Package, error) {
	var packages []structures.Package
	err := db.Joins("INNER JOIN configuration_packages as cpg ON packages.id = cpg.package_id").
		Select("packages.name, packages.version, packages.ord, packages.on_server").
		Where("cpg.rules_id = ?", rules_id).
		Find(&packages).Error
	return packages, err
}