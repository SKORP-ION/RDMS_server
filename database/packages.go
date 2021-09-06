package database

import (
	"RDMS_server/structures"
	"RDMS_server/utils"
)

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
		Joins("LEFT JOIN files ON packages.file = files.id").
		Select("packages.name, packages.version, packages.ord, packages.on_server, files.md5, cpg.action").
		Where("cpg.config_id = ?", cfg_id).
		Find(&packages).Error
	return packages, err
}

func getRulesPackagesList(rules_id uint32) ([]structures.Package, error) {
	var packages []structures.Package
	err := db.Joins("INNER JOIN configuration_packages as cpg ON packages.id = cpg.package_id").
		Joins("LEFT JOIN files ON packages.file = files.id").
		Select("packages.name, packages.version, packages.ord, packages.on_server, files.md5, cpg.action").
		Where("cpg.rules_id = ?", rules_id).
		Find(&packages).Error
	return packages, err
}

func CreateDownloadSession(pkg structures.Package) (structures.DownloadSession, error) {
	var session structures.DownloadSession


	err := db.Table("files").Joins("INNER JOIN packages as pkg ON pkg.file = files.id").
		Where("pkg.Name = ? AND pkg.Version = ?", pkg.Name, pkg.Version).
		Select("files.id as file_id, files.md5").
		Last(&session).Error

	if err != nil {
		return session, err
	}

	session.SessionKey = utils.GenerateSessionKey()

	err = db.Select("file_id", "session_key").Create(&session).Error

	return session, err
}