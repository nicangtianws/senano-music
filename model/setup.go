package model

import (
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

const DB_FILE = "Music.db"

// 初始化数据库
func InitDatabase(basedir *string) {
	dbpath := filepath.Join(*basedir, "/", DB_FILE)
	var err error
	DB, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&MusicInfo{})
	DB.AutoMigrate(&SettingInfo{})

	InitSetting(basedir)
}
