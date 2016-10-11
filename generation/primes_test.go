package generation

import (
	"math/big"
	"testing"
)

func TestMillerRabinIsPrime(t *testing.T) {
	prime := big.NewInt(137477)

	if MillerRabin(prime, 20) == false {
		t.Fatal("Number is prime, was marked as false")
	}
}

func TestMillerRabinNotPrime(t *testing.T) {
	prime := big.NewInt(7734)

	if MillerRabin(prime, 20) == true {
		t.Fatal("Number is not prime, was marked as such")
	}
}

func BenchmarkNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Naive()
	}
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Random(big.NewInt(0), big.NewInt(10))
	}
}
