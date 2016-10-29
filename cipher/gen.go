package cipher

import (
	"math/big"

	"github.com/zskamljic/rsa/gen"
	"github.com/zskamljic/rsa/primes"
)

type acceptPredicate func(r *big.Int) bool

func generateNumber(digits int, tester acceptPredicate) *big.Int {
	number := gen.RandomNDigits(digits)
	step := big.NewInt(2)
	max := big.NewInt(10)
	max.Exp(max, big.NewInt(int64(digits)), nil).Sub(max, big.NewInt(int64(1)))

	for !tester(number) {
		number.Add(number, step)

		if number.Cmp(max) > 0 {
			number = gen.RandomNDigits(digits)
		}
	}

	return number
}

func generatePrimeMiller(digits int) *big.Int {
	return generateNumber(digits, func(x *big.Int) bool { return primes.MillerRabin(x, 8) })
}
