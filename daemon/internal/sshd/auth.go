package sshd

import (
	"crypto/sha512"
	"crypto/subtle"
	"encoding/hex"
	"errors"
	"rvhc/daemon/internal/db"

	"golang.org/x/crypto/ssh"
)

func auth(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
	vmID := conn.User()

	var vm db.VM
	result := db.GetInfo(vmID, &vm)
	if result.Error != nil {
		return nil, result.Error
	}

	hashArray := sha512.Sum512(password)
	hashString := hex.EncodeToString(hashArray[:])

	if subtle.ConstantTimeCompare([]byte(vm.PwdHash), []byte(hashString)) == 1 {
		return nil, nil
	}

	return nil, errors.New("auth failed")
}
