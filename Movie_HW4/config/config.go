package config

import (
	"github.com/jinzhu/gorm"
	"main.go/structs"
)

// Koneksi ke DB
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:database12345@@tcp(127.0.0.1:3306)/homework_go_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Koneksi Database gagal!!!")
	}

	db.AutoMigrate(structs.Movie{})
	return db
}
