package main

import (
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/ssh"
)

var authPublicKeys = map[string]string{
	"user": "AAAAC3NzaC1lZDI1NTE5AAAAIADi9ZoVZstck6ELY0EIB863kD4qp5i6DYpQJHkwBiEo",
}

// publicKeyCallback handles SSH key-based authentication
// This function is largely based off of the code in this post: https://lukevers.com/2016/05/01/ssh-as-authentication-for-web-applications
func publicKeyCallback(remoteConn ssh.ConnMetadata, remoteKey ssh.PublicKey) (*ssh.Permissions, error) {
	fmt.Println("Trying to auth user " + remoteConn.User())

	// Is it a valid user?
	authPublicKey, User := authPublicKeys[remoteConn.User()]
	if !User {
		fmt.Println("User does not exist")
		return nil, errors.New("User does not exist")
	}

	authPublicKeyBytes, err := base64.StdEncoding.DecodeString(authPublicKey)
	if err != nil {
		fmt.Println("Could not base64 decode key")
		return nil, errors.New("Could not base64 decode key")
	}

	// Parse public key
	parsedAuthPublicKey, err := ssh.ParsePublicKey([]byte(authPublicKeyBytes))
	if err != nil {
		fmt.Println("Could not parse public key")
		return nil, err
	}

	// Make sure the key types match
	if remoteKey.Type() != parsedAuthPublicKey.Type() {
		fmt.Println("Key types don't match")
		return nil, errors.New("Key types do not match")
	}

	remoteKeyBytes := remoteKey.Marshal()
	authKeyBytes := parsedAuthPublicKey.Marshal()

	// Make sure the key lengths match
	if len(remoteKeyBytes) != len(authKeyBytes) {
		fmt.Println("Key lengths don't match")
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
		fmt.Println("Keys don't match")
		return nil, errors.New("Keys do not match")
	}

	return nil, nil
}
