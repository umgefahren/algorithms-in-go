package algorithms_in_go

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
)

func generateExponent() *big.Int {
	ret, err := crand.Prime(crand.Reader, 512)
	if err != nil {
		panic(err)
	}
	return ret
}

func choosePrime(e *big.Int) *big.Int {
	for {
		prime, err := crand.Prime(crand.Reader, 256)
		if err != nil {
			panic(err)
		}
		primeMin1 := big.NewInt(0)
		primeMin1.Sub(prime, big.NewInt(1))
		gg, _ := ggT(primeMin1, e)
		if gg.Cmp(big.NewInt(1)) == 0 {
			return prime
		}
		fmt.Printf("The ggt was %v \n", gg)
	}
}

func computeN(p *big.Int, q *big.Int) *big.Int {
	N := big.NewInt(0)
	N.Mul(p, q)
	return N
}

func computeEuler(p *big.Int, q *big.Int) *big.Int {
	one := big.NewInt(1)
	pMin1 := big.NewInt(0)
	pMin1.Sub(p, one)
	pMin2 := big.NewInt(0)
	pMin2.Sub(q, one)
	mul := big.NewInt(0)
	mul.Mul(pMin1, pMin2)
	return mul
}

func computeD(e *big.Int, eN *big.Int) *big.Int {
	d := big.NewInt(0)
	d.ModInverse(e, eN)
	return d
}

type PublicKey struct {
	e *big.Int
	N *big.Int
}

func (p *PublicKey) encrypt(m *big.Int) *big.Int {
	if p.N.Cmp(m) == -1 {
		panic("M is too big!")
	}
	c := big.NewInt(0)
	c.Exp(m, p.e, p.N)
	return c
}

type PrivateKey struct {
	d *big.Int
	N *big.Int
}

func (p *PrivateKey) decrypt(c *big.Int) *big.Int {
	m := big.NewInt(0)
	m.Exp(c, p.d, p.N)
	return m
}

func ComputePublicPrivateKey() (*PublicKey, *PrivateKey) {
	e := generateExponent()
	p, q := choosePrime(e), choosePrime(e)
	N := computeN(p, q)
	pubK := new(PublicKey)
	pubK.e = e
	pubK.N = N
	eN := computeEuler(p, q)
	d := computeD(e, eN)
	prvK := new(PrivateKey)
	prvK.d = d
	prvK.N = N
	return pubK, prvK
}
