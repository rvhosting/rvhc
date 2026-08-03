package db

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type VM struct {
	ID string `gorm:"primaryKey;type:varchar(36)"`
}

func (vm *VM) BeforeCreate(*gorm.DB) error {
	vm.ID = uuid.New().String()
	return nil
}
