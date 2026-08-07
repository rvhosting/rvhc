package daemon

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"rvhc/daemon/internal/dataimg"
	"rvhc/daemon/internal/db"
	"rvhc/daemon/internal/httpd"
	"rvhc/daemon/internal/sshd"
	"rvhc/daemon/internal/vmgr"
)

func Start() {
	Init()
	db.Init()
	dataimg.Init()
	vmgr.Init()
	sshd.Init(config.SSHD)
	httpd.Init()

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
