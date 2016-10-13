package main

import (
	"flag"
	"fmt"
	"math/big"

	"github.com/zskamljic/rsa/gen"
	"github.com/zskamljic/rsa/primes"
)

func main() {
	method := flag.String("method", "", "accepts \"miller\" or \"naive\" ")
	digits := flag.Int("digits", 0, "number of spaces that the generator will produce")
	bits := flag.Int("bits", 0, "number of bits that the number should generate")

	flag.Parse()

	var number *big.Int
	if *digits != 0 && *bits == 0 {
		number = gen.RandomNDigits(*digits)
	} else if *digits == 0 && *bits != 0 {
		number = gen.RandomNBits(uint(*bits))
	} else {
		if *digits == 0 && *bits == 0 {
			fmt.Println("Must specify at least one of digits or bits")
		} else {
			fmt.Println("Can't specify both digits and bits")
			return
		}
		return
	}

	fmt.Println("Generated number:")
	fmt.Println(number)

	var isPrime bool
	switch *method {
	case "miller":
		isPrime = primes.MillerRabin(number, 8)
	case "naive":
		isPrime = primes.Naive(number)
	case "":

	default:
		flag.PrintDefaults()
		return
	}

	fmt.Println("Is the number a prime?", isPrime)
}
