package dataimg

import (
	"log"
	"os"
)

var BaseDir = "dataimgs"

func Init() {
	fi, err := os.Stat(BaseDir)
	if os.IsNotExist(err) {
		goto CREATE
	}

	if err != nil {
		log.Fatalln(err)
	}

	if fi.IsDir() {
		return
	}

	if err := os.RemoveAll(BaseDir); err != nil {
		log.Fatalln(err)
	}

CREATE:
	if err := os.MkdirAll(BaseDir, 0755); err != nil {
		log.Fatalln(err)
	}
}
