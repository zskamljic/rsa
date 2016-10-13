package primes

import (
	"math/big"

	"github.com/cznic/mathutil"
	"github.com/zskamljic/rsa/gen"
)

// Naive generates a random prime number using the naive algorithm
func Naive(r *big.Int) bool {
	g := mathutil.SqrtBig(r)

	for i := big.NewInt(2); i.Cmp(g) <= 0; i.Add(i, big.NewInt(1)) {
		if big.NewInt(0).Mod(r, i).Cmp(big.NewInt(0)) == 0 {
			return false
		}
	}
	return true
}

// MillerRabin checks if the number is a prime
// returns true if r is a prime, false otherwise
func MillerRabin(r *big.Int, s int) bool {
	if r.Cmp(big.NewInt(3)) <= 0 {
		return true
	}

	cond := big.NewInt(0)

	if cond.Mod(r, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		return false
	}

	// d * 2^k = r -1
	d := big.NewInt(0).Sub(r, big.NewInt(1))
	k := big.NewInt(0)

	for cond.Mod(d, big.NewInt(2)).Cmp(big.NewInt(0)) == 0 {
		d.Div(d, big.NewInt(2))
		k.Add(k, big.NewInt(1))
	}

	// Perform test
	steps := big.NewInt(int64(s))
	genMax := big.NewInt(0).Set(r)
	genMax.Sub(genMax, big.NewInt(2))

	for j := big.NewInt(1); j.Cmp(steps) <= 0; j.Add(j, big.NewInt(1)) {
		a := gen.Random(big.NewInt(2), genMax)
		x := big.NewInt(0).Exp(a, d, r)

		if x.Cmp(big.NewInt(1)) != 0 {
			max := big.NewInt(0).Sub(k, big.NewInt(1))
			stop := big.NewInt(0).Sub(r, big.NewInt(1))

			for i := big.NewInt(0); i.Cmp(max) <= 0; i.Add(i, big.NewInt(1)) {
				if x.Cmp(stop) == 0 {
					break
				}

				x.Exp(x, big.NewInt(2), r)
			}

			if x.Cmp(stop) != 0 {
				return false
			}
		}
	}

	return true
}
