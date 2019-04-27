package main

import (
	option "bbssh/options/ssh"
	"bbssh/ssh"
	"log"
)

func main() {
	log.Println("welcome:", option.User, "\nlogin to : ", option.Host)
	ssh.SshKeyLogin(option.Host, option.User)

}
