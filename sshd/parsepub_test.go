package sshd

import (
	"encoding/base64"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"testing"
)

func TestParsePubKey(t *testing.T) {
	c, _ := ioutil.ReadFile("id_rsa.pub")
	keyBytes, err := base64.StdEncoding.DecodeString(string(c))
	if err != nil {
		log.Println("base64 decode error : ", err)
	}
	pubKey, err := ssh.ParsePublicKey(keyBytes)
	if err != nil {
		log.Println("Could not parse public key", err)
	}
	log.Println(pubKey)
}
