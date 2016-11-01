package util

import "math/big"

// ModLinEquation calculates ax≡b(mod n)
func ModLinEquation(a, n *big.Int) *big.Int {
	_, x, _ := ExtendedEuclid(a, n)

	return x.Mod(x, n)
}

// ExtendedEuclid finds the gcd of a and b -> d
// and values x and y in such a way that d = ax + by
func ExtendedEuclid(a, b *big.Int) (d, x, y *big.Int) {
	nA := big.NewInt(0).Set(a)
	nB := big.NewInt(0).Set(b)

	if b.Cmp(big.NewInt(0)) == 0 {
		d = nA
		x = big.NewInt(1)
		y = big.NewInt(0)
		return
	}

	rem := big.NewInt(0)
	rem.Mod(nA, nB)

	nD, nX, nY := ExtendedEuclid(nB, rem)

	d = nD
	x = nY
	y = big.NewInt(0)
	y.Div(nA, nB).Mul(y, nY)
	y.Sub(nX, y)

	return
}

// ModExp calculates modular exponent a^b (mod m)
func ModExp(a, b, m *big.Int) *big.Int {
	x := big.NewInt(1)
	a.Mod(a, m)

	for b.Cmp(big.NewInt(0)) != 0 {
		rem := big.NewInt(0).Set(b)
		rem.Mod(rem, big.NewInt(2))

		if rem.Cmp(big.NewInt(1)) == 0 {
			x.Mul(x, a).Mod(x, m)
		}
		b.Rsh(b, 1)
		a.Mul(a, a).Mod(a, m)
	}

	return x
}
