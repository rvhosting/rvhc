package db

import (
	"gorm.io/gorm"
)

func Delete(id string) *gorm.DB {
	return db.Where("id = ?", id).Delete(&VM{})
}
