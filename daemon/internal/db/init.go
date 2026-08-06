package db

import (
	"log"
	"sync"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var dbFile = "vm.db"
var db *gorm.DB
var mu sync.Mutex

func Init() {
	var err error
	db, err = gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.AutoMigrate(&VM{}); err != nil {
		log.Fatalln(err)
	}
}
