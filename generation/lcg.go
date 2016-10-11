package generation

import (
	"math"
	"math/big"
)

var Default = NewLcg()

type Lcg struct {
	m, a, b *big.Int
	current *big.Int
}

func NewLcg() *Lcg {
	lcg := &Lcg{
		big.NewInt(int64(math.Pow(2, 32))),
		big.NewInt(69069),
		big.NewInt(0),
		big.NewInt(1),
	}
	return lcg
}

func (this *Lcg) Next() *big.Int {
	this.current = this.current.
		Mul(this.a, this.current).
		Add(this.current, this.b).
		Mod(this.current, this.m)

	return this.current
}

func Random(a, b *big.Int) *big.Int {
	mod := big.NewInt(0).Sub(b, a)
	mod.Add(mod, big.NewInt(1))

	random := Default.Next()

	rhs := big.NewInt(0).Mod(random, mod)

	return a.Add(a, rhs)
}
