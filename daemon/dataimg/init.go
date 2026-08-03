package dataimg

import (
	"log"
	"os"
)

var baseDir = "dataimgs"

func Init() {
	fi, err := os.Stat(baseDir)
	if os.IsNotExist(err) {
		goto CREATE
	}

	if err != nil {
		log.Fatalln(err)
	}

	if fi.IsDir() {
		return
	}

	if err := os.RemoveAll(baseDir); err != nil {
		log.Fatalln(err)
	}

CREATE:
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		log.Fatalln(err)
	}
}
