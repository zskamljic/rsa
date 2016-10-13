package gen

import (
	"math"
	"math/big"
	"testing"
)

func TestGenerateNDigits(t *testing.T) {
	for i := 1; i < 10; i++ {
		num := RandomNDigits(i)

		max := int64(math.Pow(10, float64(i)))
		mod := num.Mod(num, big.NewInt(max))

		if mod.Cmp(num) != 0 {
			t.Fatal("Generated number did not have", i, "digits: ", num)
		}
	}
}

func TestGenerateNBits(t *testing.T) {
	for i := 1; i < 10; i++ {
		num := RandomNBits(uint(i))

		if num.BitLen() != i {
			t.Fatal("Generated number did not have", i, "bits: ", num.BitLen())
		}
	}
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Random(big.NewInt(0), big.NewInt(10))
	}
}
