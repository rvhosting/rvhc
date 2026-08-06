package db

import (
	"gorm.io/gorm"
)

func GetIDs(dest any) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()

	return db.Model(&VM{}).Pluck("id", dest)
}
