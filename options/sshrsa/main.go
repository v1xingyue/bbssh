package sshrsa

import (
	"flag"
	"os"
)

var (
	Name     = "sshrsa option"
	BindAddr = ""
	User     = ""
	Help     = false
	KeyDir   = ".rsa"
)

func init() {
	flag.StringVar(&BindAddr, "b", "localhost:12034", "Bind Addr.")
	flag.StringVar(&User, "u", "", "Generate one pair RSA keys.")
	flag.StringVar(&KeyDir, "kdir", ".rsa", "Dir you want to save rsa private keys.")
	flag.BoolVar(&Help, "h", false, "Help ")

	flag.Parse()
	if Help {
		flag.Usage()
		os.Exit(0)
	}
}
