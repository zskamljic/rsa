package main

import (
	"fmt"
	"math/big"

	"github.com/zskamljic/rsa/generation"
)

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println(generation.Naive())
	}

	fmt.Println()

	for i := 0; i < 15; i++ {
		fmt.Println(generation.Random(big.NewInt(0), big.NewInt(10)))
	}
}
