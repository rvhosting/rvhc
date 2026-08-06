package sshd

import (
	"context"
	"log/slog"
	"net"

	"github.com/cmoog/sshproxy"
	"golang.org/x/crypto/ssh"
)

func process(conn net.Conn, server ssh.ServerConfig) {
	defer conn.Close()

	sshConn, sshChannels, sshRequests, err := ssh.NewServerConn(conn, &server)
	if err != nil {
		slog.Error(err.Error())
		return
	}

	targetAddr, targetClientConfig, err := routing(sshConn.User())
	if err != nil {
		slog.Error(err.Error())
		return
	}

	proxy := sshproxy.New(targetAddr, targetClientConfig)
	proxy.Serve(context.Background(), sshConn, sshChannels, sshRequests)
}
