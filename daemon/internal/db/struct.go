package db

import (
	"github.com/google/uuid"

	"gorm.io/gorm"
)

type VM struct {
	ID       string `gorm:"primaryKey;type:varchar(36);not null" json:"id"`
	CPU      uint8  `gorm:"not null" json:"cpu"`
	Memory   uint16 `gorm:"not null" json:"memory"`    // MB
	DataSize uint16 `gorm:"not null" json:"data_size"` // MB
	Running  bool   `gorm:"not null" json:"running"`
	PwdHash  string `gorm:"not null" json:"pwd_hash"`
}

func (vm *VM) BeforeCreate(*gorm.DB) error {
	vm.ID = uuid.New().String()
	return nil
}
