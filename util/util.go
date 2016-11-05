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
	aParam := big.NewInt(0).Set(a)
	bParam := big.NewInt(0).Set(b)
	aParam.Mod(aParam, m)

	for bParam.Cmp(big.NewInt(0)) != 0 {
		rem := big.NewInt(0).Set(bParam)
		rem.Mod(rem, big.NewInt(2))

		if rem.Cmp(big.NewInt(1)) == 0 {
			x.Mul(x, aParam).Mod(x, m)
		}
		bParam.Rsh(bParam, 1)
		aParam.Mul(aParam, aParam).Mod(aParam, m)
	}

	return x
}
