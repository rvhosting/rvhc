package downloader

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var baseDir = "minihttpsys"
var baseUrl = "https://github.com/rvhosting/minihttpsys/releases/latest/download"
var files = []string{"bios.bin", "kernel.bin", "sys.img"}

func Download() {
	cleandir()

	for _, file := range files {
		path := filepath.Join(baseDir, file)
		url := fmt.Sprintf("%s/%s", baseUrl, file)

		log.Println("download", url, "to", path)
		func() {
			f, err := os.Create(path)
			if err != nil {
				log.Fatalln(err)
			}

			resp, err := http.Get(url)
			if err != nil {
				log.Fatalln(err)
			}

			defer f.Close()
			defer resp.Body.Close()
			io.Copy(f, resp.Body)
		}()
	}
}
