package algorithms_in_go

import (
	"math/big"
	"testing"
)

func TestRsa(t *testing.T) {
	pub, prv := ComputePublicPrivateKey()
	m := big.NewInt(0)
	c := pub.encrypt(m)
	mc := prv.decrypt(c)
	if mc.Cmp(m) != 0 {
		t.Errorf("The comparison failed. The m was %v and the decrypted was %v", m, mc)
	}
}
