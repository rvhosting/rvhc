package daemon

import (
	"encoding/json"
	"log"
	"os"
)

var ConfigFile string
var config daemonConfig

func Init() {
	f, err := os.Open(ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&config); err != nil {
		log.Fatalln(err)
	}
}
