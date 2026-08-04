package db

import (
	"gorm.io/gorm"
)

func GetIDs(dest any) *gorm.DB {
	return db.Model(&VM{}).Pluck("id", dest)
}
