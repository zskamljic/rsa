package generation

import (
	"math/big"
	"testing"
)

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Random(big.NewInt(0), big.NewInt(10))
	}
}
