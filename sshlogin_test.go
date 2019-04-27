package bbssh

import (
	"bbssh/ssh"
	"testing"
)

func TestLoadSshKey(t *testing.T) {
	t.Log("private key : ", ssh.LoadSshPriKey("xingyue", "10.210.12.18"))
}

//func TestLoginWithKey(t *testing.T) {
//	user := "root"
//	keystr := `
//-----BEGIN RSA PRIVATE KEY-----
//TVVyMs8AQ4yfpkfaYzNCWTx6Z1Fu1tWWiw5lFPVj+qgEPiHw64D7pWOVXYZTnqtl
//noJ1nnLKai2GZ67okRd55HkoYNdRcF2YjHK4ldVeQradt4AxVlTs
//.........
//-----END RSA PRIVATE KEY-----
//	`
//	host := "1.1.1.1:22"
//	loginWithKey(keystr, host, user)
//}
//
//func TestLoginWithPassword(t *testing.T) {
//	user := "v1xingyue"
//	password := ""
//	host := "1.1.1.1:22"
//	loginWithPassword(password, host, user)
//}
