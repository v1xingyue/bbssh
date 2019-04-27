package main

import (
	_ "bbssh/options/sshd"
	"bbssh/sshd"
)

func main() {
	sshd.StartSshd()
}
