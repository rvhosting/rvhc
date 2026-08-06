package sshd

import (
	"fmt"
	"rvhc/daemon/internal/vmgr"

	"golang.org/x/crypto/ssh"
)

func routing(vmID string) (string, *ssh.ClientConfig, error) {
	portMap, err := vmgr.GetPorts(vmID)
	if err != nil {
		return "", nil, err
	}

	addr := fmt.Sprintf("127.0.0.1:%d", portMap[22])
	clientConfig := &ssh.ClientConfig{
		User:            "root",
		Auth:            []ssh.AuthMethod{ssh.Password("")},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	return addr, clientConfig, nil
}
