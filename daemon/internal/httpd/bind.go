package httpd

import (
	"encoding/json"
	"os"
)

func Bind(host string, id string) error {
	mu.Lock()
	defer mu.Unlock()

	hosts[host] = id
	f, err := os.OpenFile(hostsFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	return json.NewEncoder(f).Encode(hosts)
}
