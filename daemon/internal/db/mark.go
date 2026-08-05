package db

import (
	"gorm.io/gorm"
)

func Mark(id string, running bool) *gorm.DB {
	return db.Model(&VM{}).Where("id = ?", id).Update("running", running)
}
