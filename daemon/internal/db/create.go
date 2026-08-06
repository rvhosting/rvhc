package db

import (
	"gorm.io/gorm"
)

func Create(vm *VM) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()

	return db.Create(vm)
}
