package downloader

import (
	"log"
	"os"
)

func cleandir() {
	if err := os.RemoveAll(baseDir); err != nil {
		log.Fatalln(err)
	}

	if err := os.MkdirAll(baseDir, 0755); err != nil {
		log.Fatalln(err)
	}
}
