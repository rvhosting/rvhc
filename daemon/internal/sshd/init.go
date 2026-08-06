package sshd

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"path/filepath"

	"github.com/fasmide/hostkeys"
	"golang.org/x/crypto/ssh"
)

var hostKeyDir = "sshd_host_key"

func Init(port int) {
	if err := os.MkdirAll(hostKeyDir, 0755); err != nil {
		log.Fatalln(err)
	}

	server := ssh.ServerConfig{
		PasswordCallback: auth,
	}

	manager := &hostkeys.Manager{
		Directory:    filepath.Join(".", hostKeyDir),
		NamingScheme: "ssh_host_%s_key",
	}

	if err := manager.Manage(&server); err != nil {
		log.Fatalln(err)
	}

	addr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("start sshd")
	go func(listener net.Listener, server ssh.ServerConfig) {
		for {
			conn, err := listener.Accept()
			if err != nil {
				slog.Error(err.Error())
				continue
			}

			go process(conn, server)
		}
	}(listener, server)
}
