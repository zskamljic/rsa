package gen

import (
	"math"
	"math/big"
	"time"
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

	for i := time.Now().Nanosecond() % 10000; i >= 0; i-- {
		lcg.Next()
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
