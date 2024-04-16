package database

func migrate() {
	DbClient.AutoMigrate(&File{})
}
