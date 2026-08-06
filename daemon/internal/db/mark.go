package db

import (
	"gorm.io/gorm"
)

func Mark(id string, running bool) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()

	return db.Model(&VM{}).Where("id = ?", id).Update("running", running)
}
