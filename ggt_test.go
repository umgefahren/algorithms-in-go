package algorithms_in_go

import (
	"math/big"
	"testing"
)

func TestGGT(t *testing.T) {
	n11 := big.NewInt(1736828125)
	n12 := big.NewInt(683826495)
	correct1 := big.NewInt(5)
	res11, _ := ggT(n11, n12)
	if res11.Cmp(correct1) != 0 {
		t.Errorf("The ggt of %v and %v is %v not %v", n11, n12, correct1, res11)
	}
	res12, _ := ggT(n12, n11)
	if res12.Cmp(correct1) != 0 {
		t.Errorf("The ggt of %v and %v is %v not %v", n12, n11, correct1, res12)
	}
}
