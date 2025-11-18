package model

import "gorm.io/gorm"

type SettingInfo struct {
	gorm.Model
	Id    int    `json:"id" gorm:"unique;primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

func InitSetting(basedir *string) {
	settingInfo := SettingInfo{
		Name:  "重新扫描",
		Value: "扫描",
		Type:  "BUTTON",
	}
	DB.Create(&settingInfo)
}
