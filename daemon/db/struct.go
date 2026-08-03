package db

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type VM struct {
	ID      string `gorm:"primaryKey;type:varchar(36);not null" json:"id"`
	DataImg string `gorm:"type:text;not null;unique" json:"data_img"`
	CPU     int    `gorm:"not null" json:"cpu"`
	Memory  int    `gorm:"not null" json:"memory"`
}

func (vm *VM) BeforeCreate(*gorm.DB) error {
	vm.ID = uuid.New().String()
	return nil
}
