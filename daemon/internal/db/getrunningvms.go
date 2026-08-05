package db

import (
	"gorm.io/gorm"
)

func GetRunningVMs(dest any) *gorm.DB {
	return db.Where("running = ?", true).Find(&dest)
}
