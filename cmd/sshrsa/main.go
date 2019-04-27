package main

import (
	option "bbssh/options/sshrsa"
	"bbssh/sshrsa"
	"bbssh/sshrsa/api"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	if option.User == "" {
		api.StartServer(option.BindAddr)
	} else {
		pri, pub, err := sshrsa.MakeSSHKeyPair(2048)
		fpath := fmt.Sprintf("%s/%s_rsa.pri", option.KeyDir, option.User)
		fpub := fmt.Sprintf("%s/%s_rsa.pub", option.KeyDir, option.User)
		if err == nil {
			err = ioutil.WriteFile(fpath, []byte(pri), 0600)
			err = ioutil.WriteFile(fpub, []byte(pub), 0600)

		}
		if err == nil {
			log.Println("generate ok!")
		} else {
			log.Println("error happend!", err)
		}
	}
}
