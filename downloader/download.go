package downloader

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
)

var baseDir = "minihttpsys"
var baseUrl = "https://github.com/rvhosting/minihttpsys/releases/latest/download"
var files = []string{"bios.bin", "kernel.bin", "sys.img"}
var MaxRetry int

func Download() {
	cleandir()

	for _, file := range files {
		path := filepath.Join(baseDir, file)
		url := fmt.Sprintf("%s/%s", baseUrl, file)

		log.Println("download", url, "to", path)

		for i := 1; i <= MaxRetry; i++ {
			err := func() error {
				f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
				if err != nil {
					return err
				}
				defer f.Close()

				resp, err := http.Get(url)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				if _, err := io.Copy(f, resp.Body); err != nil {
					return err
				}

				return nil
			}()

			if err != nil {
				if i == MaxRetry {
					log.Fatalln(err)
				}

				slog.Error(err.Error(), "retry", i)
			}
		}
	}
}
