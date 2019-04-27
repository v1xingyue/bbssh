package sshrsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
)

func GenerateKey(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	private, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return private, &private.PublicKey, nil
}

func EncodePrivateKey(private *rsa.PrivateKey) []byte {
	return pem.EncodeToMemory(&pem.Block{
		Bytes: x509.MarshalPKCS1PrivateKey(private),
		Type:  "RSA PRIVATE KEY",
	})
}

func EncodeSSHKey(public *rsa.PublicKey) ([]byte, error) {
	publicKey, err := ssh.NewPublicKey(public)
	if err != nil {
		return nil, err
	}
	return ssh.MarshalAuthorizedKey(publicKey), nil
}

// 默认ssh rsa 长度为2048
func MakeSSHKeyPair(bits int) (string, string, error) {
	pkey, pubkey, err := GenerateKey(bits)
	if err != nil {
		return "", "", err
	}
	pub, err := EncodeSSHKey(pubkey)
	if err != nil {
		return "", "", err
	}
	return string(EncodePrivateKey(pkey)), string(pub), nil
}
