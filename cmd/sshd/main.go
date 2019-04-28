package main

import (
	_ "bbssh/options/sshd"
	"bbssh/sshd"
	"bbssh/user"
	"log"
)

func main() {
	if user.IsRoot() {
		sshd.StartSshd()
	} else {
		log.Fatal("You Must Run This Under Root User")
	}
}
