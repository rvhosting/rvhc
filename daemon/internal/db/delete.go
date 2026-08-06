package db

import (
	"gorm.io/gorm"
)

func Delete(id string) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()

	return db.Where("id = ?", id).Delete(&VM{})
}
