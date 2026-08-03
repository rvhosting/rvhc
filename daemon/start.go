package daemon

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"rvhc/daemon/db"
)

var ConfigFile string
var config daemonConfig

func Start() {
	db.Init()

	f, err := os.Open(ConfigFile)
	if err != nil {
		log.Fatalln(err)
	}

	if err := json.NewDecoder(f).Decode(&config); err != nil {
		log.Fatalln(err)
	}

	addr := fmt.Sprintf(
		"%s:%d",
		config.Listen.Host,
		config.Listen.Port,
	)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("start daemon")
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error(err.Error())
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()

			if !auth(conn) {
				return
			}

			process(conn)
		}(conn)
	}
}
