package db

import (
	"gorm.io/gorm"
)

func GetInfo(id string, dest any) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()

	return db.Where("id = ?", id).First(dest)
}
