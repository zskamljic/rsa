package cipher

import (
	"fmt"
	"math/big"
)

// Cipher is used to encipher or decipher RSA
type Cipher struct {
}

// NewCipher returns a new Cipher
func NewCipher() *Cipher {
	return &Cipher{}
}

// GenerateCipher generates a cipher with p and q
// containing @digits digits.
func GenerateCipher(digits int) *Cipher {
	var p, q *big.Int

	for {
		p = generatePrimeMiller(digits)
		q = generatePrimeMiller(digits)

		if p.Cmp(q) != 0 {
			break
		}
	}

	n := big.NewInt(0)
	n.Mul(p, q)

	p.Sub(p, big.NewInt(1))
	q.Sub(q, big.NewInt(1))

	euler := big.NewInt(0)
	euler.Mul(p, q)

	var e *big.Int
	gcd := big.NewInt(0)
	for {
		e = generatePrimeMiller(digits - 1)

		if gcd.GCD(nil, nil, e, euler).Cmp(big.NewInt(1)) == 0 {
			break
		}
	}

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(n)
	fmt.Println(e)
	fmt.Println(euler)

	return NewCipher()
}
