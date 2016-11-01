package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/zskamljic/rsa/cipher"
)

func main() {
	decode := flag.Bool("decode", false, "")
	digits := flag.Int("digits", 10, "the number of digits for p and q")

	flag.Parse()

	scanner := bufio.NewReader(os.Stdin)

	if *decode {
		fmt.Println("Decode!")
	} else {
		fmt.Print("Message to encode: ")
		message, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Println(message)

		cipher := cipher.GenerateCipher(*digits)
		cipher.SaveKeys()

		encoded := cipher.Encode(message)

		fmt.Println(encoded)
		fmt.Println(string(cipher.Decode(encoded)))
	}
}
