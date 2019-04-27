package sshrsa

import (
	"flag"
	"os"
)

var (
	Name     = "sshd option"
	BindAddr = ""
	RsaFile  = "id_rsa"
	Help     = false
)

func init() {
	flag.StringVar(&BindAddr, "b", "localhost:12099", "Bind Addr.")
	flag.StringVar(&RsaFile, "f", "id_rsa", "id_rsa sshd use local.")
	flag.BoolVar(&Help, "h", false, "Help ")

	flag.Parse()
	if Help {
		flag.Usage()
		os.Exit(0)
	}
}
