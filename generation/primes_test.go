package generation

import (
	"math/big"
	"testing"
)

func TestMillerRabinFirst1000Primes(t *testing.T) {

	primes := 0
	for i := 2; i < 1000; i++ {
		if MillerRabin(big.NewInt(int64(i)), 20) {
			primes++
		}
	}

	if primes != 168 {
		t.Fatal("Expected number of primes: 168, actual:", primes)
	}
}

func TestMillerRabinPrimeReturnsTrue(t *testing.T) {
	prime := big.NewInt(100003)

	if MillerRabin(prime, 20) == false {
		t.Fatal("Number is prime, algorithm returned false")
	}
}

func TestMillerRabinNotPrimeReturnsFalse(t *testing.T) {
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
