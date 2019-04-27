package ssh

import (
	"flag"
	"os"
)

var (
	Name    = "ssh option"
	KeyHost = "localhost:12034"
	Host    = ""
	User    = ""
	Help    = false
)

func init() {
	flag.StringVar(&KeyHost, "khost", "localhost:12034", "host get private key.")
	flag.StringVar(&Host, "host", "", "host want to connect.")
	flag.StringVar(&User, "user", "", "user login.")
	flag.BoolVar(&Help, "h", false, "Help ")

	flag.Parse()
	if Help {
		flag.Usage()
		os.Exit(0)
	}
}
