package db

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type VM struct {
	ID      string `gorm:"primaryKey;type:varchar(36)" json:"id"`
	DataImg string `gorm:"type:text;not null;unique" json:"data"`
	CPU     uint8  `gorm:"not null" json:"cpu"`
	Memory  uint16 `gorm:"not null" json:"memory"` // MB
}

func (vm *VM) BeforeCreate(*gorm.DB) error {
	vm.ID = uuid.New().String()
	return nil
}
