package database

import "gorm.io/gorm/clause"

func AddFilesToDb(files []File) {
	DbClient.Create(&files)
}

func GetFilesByBasePath(directory string) []File {
	var files []File
	DbClient.Where("path LIKE ?", directory+"%").Order(clause.OrderByColumn{
		Column: clause.Column{
			Name: "modified_at",
		},
		Desc: true,
	}).Find(&files)
	return files
}
