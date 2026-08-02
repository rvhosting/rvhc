package daemon

import (
	"encoding/json"
	"log"
	"os"
)

var ConfigFile string
var config daemonConfig

func Start() {
	f, err := os.Open(ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.NewDecoder(f).Decode(&config); err != nil {
		log.Fatalln(err)
	}
}
