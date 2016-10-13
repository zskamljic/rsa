package generation

import (
	"math"
	"math/big"
)

// Default Lcg
var Default = defaultLcg()

// Lcg is the generator struct
type Lcg struct {
	m, a, b *big.Int
	current *big.Int
}

func defaultLcg() *Lcg {
	return NewLcg(int64(math.Pow(2, 32)), 69069, 0, 1)
}

// NewLcg creates a new generator using the given parameters
func NewLcg(m, a, b, rn0 int64) *Lcg {
	lcg := &Lcg{
		big.NewInt(m),
		big.NewInt(a),
		big.NewInt(b),
		big.NewInt(rn0),
	}
	return lcg
}

// Next returns the next number in the sequence
func (lcg *Lcg) Next() *big.Int {
	lcg.current = lcg.current.
		Mul(lcg.a, lcg.current).
		Add(lcg.current, lcg.b).
		Mod(lcg.current, lcg.m)

	return lcg.current
}

// Random returns a random number in range a, b
func Random(a, b *big.Int) *big.Int {
	mod := big.NewInt(0).Sub(b, a)
	mod.Add(mod, big.NewInt(1))

	random := Default.Next()

	rhs := big.NewInt(0).Mod(random, mod)

	return a.Add(a, rhs)
}
