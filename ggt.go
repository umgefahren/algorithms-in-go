package algorithms_in_go

import "math/big"

func ggT(a *big.Int, b *big.Int) (*big.Int, *big.Int) {
	zero := big.NewInt(0)

	if b.Cmp(zero) == 0 {
		return a, b
	}

	c := big.NewInt(0)
	c = c.Mod(a, b)

	return ggT(b, c)
}
