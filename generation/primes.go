package generation

import (
	"math/big"

	"github.com/cznic/mathutil"
)

func Naive() *big.Int {
generate:
	for {
		r := Default.Next()

		g := mathutil.SqrtBig(r)

		step := big.NewInt(2)
		for i := big.NewInt(3); i.Cmp(g) <= 0; i.Add(i, step) {
			if big.NewInt(0).Mod(r, i) == big.NewInt(0) {
				continue generate
			}
		}
		return r
	}
}

func MillerRabin(r *big.Int, s int) bool {
	if r.Cmp(big.NewInt(3)) <= 0 {
		return true
	}

	rem := big.NewInt(0)
	if rem.Mod(r, big.NewInt(2)) == big.NewInt(0) {
		return false
	}

	d := big.NewInt(0).Sub(r, big.NewInt(1))
	k := big.NewInt(0)

	for rem.Mod(d, big.NewInt(2)) == big.NewInt(0) {
		d.Div(d, big.NewInt(2))
		k.Add(k, big.NewInt(1))
	}

	genMax := big.NewInt(0).Set(r)
	genMax.Sub(genMax, big.NewInt(2))
	for j := big.NewInt(1); j.Cmp(big.NewInt(int64(s))) < 0; j.Add(j, big.NewInt(1)) {
		a := Random(big.NewInt(2), genMax)
		x := a.Exp(a, d, nil)

		if x != big.NewInt(1) {
			max := big.NewInt(0).Sub(k, big.NewInt(1))
			stop := big.NewInt(0)
			stop.Sub(r, big.NewInt(1))

			for i := big.NewInt(0); i.Cmp(max) < 0; i.Add(i, big.NewInt(1)) {
				if x == stop {
					break
				}

				x = x.Exp(x, big.NewInt(2), nil)
				x.Mod(x, r)
			}

			if x.Cmp(stop) != 0 {
				return false
			}
		}
	}

	return true
}
