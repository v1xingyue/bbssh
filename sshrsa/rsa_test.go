package sshrsa

import (
	"testing"
)

func TestRsaPair(t *testing.T) {
	pkey, pubkey, _ := GenerateKey(2048)
	pub, _ := EncodeSSHKey(pubkey)
	t.Log("----- public -------- ")
	t.Log(string(pub))
	t.Log("----- private -------- ")
	t.Log(string(EncodePrivateKey(pkey)))
	t.Log(MakeSSHKeyPair(2048))
}
