package sshd

import (
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/ssh"
	"log"
)

var authPublicKeys = map[string]string{
	"xingyue": "AAAAB3NzaC1yc2EAAAADAQABAAABAQCveKzQBwRyA9cWoA/810hx3OEmONt4D38CBIDxnlV9u2bxy7Y7KIW9VX2cxX4d7XQJj87nJTjiMivg0pkcNLC+4RHby/cPV5rSv1n4qbOAPxVJGhg9udeD80EcjbKSqOGlO2WPRGbNhN9UM0YA+WXTRwWlUZxrEydE55D0C5s/NbfWYh4dqUawgh58nKrjbeix3g+r+lhEzOxW0DE9lbMYIt5WsGojvf/NZV+VNY5q+PNXTm9xzNzfTB0c60pY+677hECL0qlu5M5gyG4Pskz8pTb9tvUkaHOf+Q8y7CV0AX1JEifu/xQoc57rJ5/Rk5N0zGN5xCTMvMryTtySHlkR",
}

func publicKeyCallback(remoteConn ssh.ConnMetadata, remoteKey ssh.PublicKey) (*ssh.Permissions, error) {
	log.Println("Trying to auth user " + remoteConn.User())

	// Is it a valid user?
	userPubKeys, ok := authPublicKeys[remoteConn.User()]
	if !ok {
		log.Println("User does not exist")
		return nil, errors.New("User does not exist")
	}

	authPublicKeyBytes, err := base64.StdEncoding.DecodeString(string(userPubKeys))
	if err != nil {
		log.Println("Could not base64 decode key")
		return nil, errors.New("Could not base64 decode key")
	}

	// Parse public key
	parsedAuthPublicKey, err := ssh.ParsePublicKey([]byte(authPublicKeyBytes))
	if err != nil {
		log.Println("Could not parse public key")
		return nil, err
	}

	// Make sure the key types match
	if remoteKey.Type() != parsedAuthPublicKey.Type() {
		log.Println("Key types don't match")
		return nil, errors.New("Key types do not match")
	}

	remoteKeyBytes := remoteKey.Marshal()
	authKeyBytes := parsedAuthPublicKey.Marshal()

	// Make sure the key lengths match
	if len(remoteKeyBytes) != len(authKeyBytes) {
		log.Println("Key lengths don't match")
		return nil, errors.New("Keys do not match")
	}

	// Make sure every byte of the key matches up
	// TODO: This should be a constant time check
	keysMatch := true
	for i, b := range remoteKeyBytes {
		if b != authKeyBytes[i] {
			keysMatch = false
		}
	}

	if keysMatch == false {
		log.Println("Keys don't match")
		return nil, errors.New("Keys do not match")
	}

	return nil, nil
}
