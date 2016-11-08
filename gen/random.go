package gen

import "math/big"

// Random returns a random number in range a, b
func Random(a, b *big.Int) *big.Int {
	mod := big.NewInt(0).Sub(b, a)
	mod.Add(mod, big.NewInt(1))
	if mod.Cmp(big.NewInt(0)) == 0 {
		mod.SetInt64(1)
	}

	random := Default.Next()

	rhs := big.NewInt(0).Mod(random, mod)

	return a.Add(a, rhs)
}

// RandomNDigits returns a random number with n digits
func RandomNDigits(digits int) *big.Int {
	min := big.NewInt(10)
	max := big.NewInt(10)

	min.Exp(min, big.NewInt(int64(digits-1)), nil)
	max.Exp(max, big.NewInt(int64(digits)), nil)
	max.Sub(max, big.NewInt(1))

	return Random(min, max)
}

// RandomNBits generates a random number that is n bits long
func RandomNBits(bits uint) *big.Int {
	if bits == 0 {
		return big.NewInt(0)
	} else if bits == 1 {
		return Random(big.NewInt(0), big.NewInt(1))
	}

	min := big.NewInt(1)
	min.Lsh(min, bits-1).Sub(min, big.NewInt(1))
	max := big.NewInt(1)
	max.Lsh(max, bits).Sub(max, big.NewInt(1))

	random := Random(min, max)

	mask := big.NewInt(1)
	for i := uint(0); i < bits; i++ {
		mask.Lsh(mask, 1).Or(mask, big.NewInt(1))
	}

	result := big.NewInt(0)
	result.And(random, mask)
	return result.SetBit(result, int(bits-1), 1)
}
