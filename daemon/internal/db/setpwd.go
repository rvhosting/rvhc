package db

import (
	"crypto/sha512"
	"encoding/hex"

	"gorm.io/gorm"
)

func SetPassword(id string, pwd string) *gorm.DB {
	hashArray := sha512.Sum512([]byte(pwd))
	hashString := hex.EncodeToString(hashArray[:])
	return db.Model(&VM{}).Where("id = ?", id).Update("pwd_hash", hashString)
}
