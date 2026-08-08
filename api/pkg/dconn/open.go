package dconn

import (
	"fmt"
	"log"
	"net"
	"rvhc/api/pkg/config"
	"rvhc/protocol"
)

var conn net.Conn

func Open(dc config.DaemonConfig) {
	addr := net.JoinHostPort(
		dc.Listen.Host,
		fmt.Sprintf("%d", dc.Listen.Port),
	)

	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	if err := protocol.Write(conn, dc.Auth); err != nil {
		log.Fatalln(err)
	}

	var resp protocol.Response
	if err := protocol.Read(conn, &resp); err != nil {
		log.Fatalln(err)
	}

	if !resp.Success {
		log.Fatalln(resp.Message)
	}

	conn = c
}
