package db

import (
	"gorm.io/gorm"
)

func GetInfo(id string, dest any) *gorm.DB {
	return db.Where("id = ?", id).First(dest)
}
