package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dsn := "link-db-user:link-db-user@tcp(mysql)/link-ses-db?parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("DB接続に失敗しました")
	}
	return db
}
