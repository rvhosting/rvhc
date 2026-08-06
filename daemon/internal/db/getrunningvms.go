package db

import (
	"gorm.io/gorm"
)

func GetRunningVMs(dest any) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()

	return db.Where("running = ?", true).Find(dest)
}
