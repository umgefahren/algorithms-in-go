package algorithms_in_go

import (
	"crypto/rand"
	"math/big"
)

func generatePrime() *big.Int {
	p, err := rand.Prime(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return p
}

func generateBase() *big.Int {
	g := big.NewInt(1000)
	g, err := rand.Int(rand.Reader, g)
	if err != nil {
		panic(err)
	}
	return g
}

type SharedParams struct {
	p *big.Int
	g *big.Int
}

func NewSharedParams() *SharedParams {
	sp := new(SharedParams)
	sp.p = generatePrime()
	sp.g = generateBase()
	return sp
}

type Secret struct {
	x  *big.Int
	sp *SharedParams
}

func NewSecret(sp *SharedParams) *Secret {
	pMin2 := big.NewInt(0)
	pMin2.Sub(sp.p, big.NewInt(2))
	xA, err := rand.Int(rand.Reader, pMin2)
	if err != nil {
		panic(err)
	}
	ret := new(Secret)
	ret.x = xA
	ret.sp = sp
	return ret
}

func (s *Secret) computeY() *big.Int {
	y := big.NewInt(0)
	y.Exp(s.sp.g, s.x, s.sp.p)
	return y
}

func (s *Secret) computeShared(y *big.Int) *big.Int {
	k := big.NewInt(0)
	k.Exp(y, s.x, s.sp.p)
	return k
}
