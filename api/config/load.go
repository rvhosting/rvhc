package config

import (
	"encoding/json"
	"log"
	"os"
)

var Daemon DaemonConfig

func Load(file, name string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	daemons := map[string]DaemonConfig{}
	if err := json.NewDecoder(f).Decode(&daemons); err != nil {
		log.Fatalln(err)
	}

	d, exist := daemons[name]
	if !exist {
		log.Fatalln("unknown daemon")
	}

	Daemon = d
}
