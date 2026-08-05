package downloader

import (
	"log"
	"os"
)

func cleandir() {
	if err := os.RemoveAll(BaseDir); err != nil {
		log.Fatalln(err)
	}

	if err := os.MkdirAll(BaseDir, 0755); err != nil {
		log.Fatalln(err)
	}
}
