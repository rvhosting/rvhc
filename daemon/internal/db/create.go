package db

import (
	"gorm.io/gorm"
)

func Create(vm *VM) *gorm.DB {
	return db.Create(vm)
}
