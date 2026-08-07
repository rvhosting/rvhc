package httpd

import (
	"encoding/json"
	"os"
)

func LoadHosts() error {
	mu.Lock()
	defer mu.Unlock()

	fi, err := os.Stat(hostsFile)
	if os.IsNotExist(err) {
		goto CREATE
	}

	if err != nil {
		return err
	}

	if fi.IsDir() {
		if err := os.RemoveAll(hostsFile); err != nil {
			return err
		}

	} else {
		goto LOAD
	}

CREATE:
	{
		f, err := os.Create(hostsFile)
		if err != nil {
			return err
		}

		if _, err := f.WriteString("{}"); err != nil {
			return err
		}
	}

LOAD:
	{
		f, err := os.Open(hostsFile)
		if err != nil {
			return err
		}

		if err := json.NewDecoder(f).Decode(&hosts); err != nil {
			return err
		}
	}

	return nil
}
