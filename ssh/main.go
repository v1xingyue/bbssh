package ssh

import (
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"net"
	"os"
	"time"
)

const (
	ModePassword int = iota
	ModeKey      int = iota
)

func nilCallBack(hostname string, remote net.Addr, key ssh.PublicKey) error {
	return nil
}

func errorCallBack(err error, msg string) {
	if err != nil {
		log.Fatalf("%s error: %v", msg, err)
	}
}

func runSession(session *ssh.Session) {
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.ECHOCTL:       0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	termFD := int(os.Stdin.Fd())
	w, h, _ := terminal.GetSize(termFD)
	termState, _ := terminal.MakeRaw(termFD)
	defer terminal.Restore(termFD, termState)
	err := session.RequestPty("xterm-256color", h, w, modes)
	errorCallBack(err, "request pty")
	err = session.Shell()
	errorCallBack(err, "start shell")
	err = session.Wait()
	//errorCallBack(err, "return")
}

func makeAuth(key string, t int) []ssh.AuthMethod {
	auth := make([]ssh.AuthMethod, 0)
	if t == ModePassword {
		auth = append(auth, ssh.Password(key))
	} else {
		signer, err := ssh.ParsePrivateKey([]byte(key))
		if err != nil {
			log.Fatal(err)
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}
	return auth
}

func loginWithPassword(password string, host string, user string) {
	client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: nilCallBack,
		Timeout:         30 * time.Second,
	})
	errorCallBack(err, "dial")
	session, err := client.NewSession()
	errorCallBack(err, "client new session")
	runSession(session)
}

func loginWithKey(keystr string, host string, user string) {
	auth := make([]ssh.AuthMethod, 0)
	signer, err := ssh.ParsePrivateKey([]byte(keystr))
	if err != nil {
		log.Fatal(err)
	}
	auth = append(auth, ssh.PublicKeys(signer))

	client, err := ssh.Dial("tcp", host, &ssh.ClientConfig{
		User:            user,
		Auth:            auth,
		HostKeyCallback: nilCallBack,
		Timeout:         2 * time.Second,
	})
	errorCallBack(err, "dial")
	session, err := client.NewSession()
	errorCallBack(err, "client new session")
	runSession(session)
}

func SshKeyLogin(host string, user string) {
	pri := LoadSshPriKey(user, host)
	loginWithKey(pri, host, user)
}
