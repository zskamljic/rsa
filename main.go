package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"

	"io/ioutil"

	"github.com/zskamljic/rsa/cipher"
)

func main() {
	decode := flag.Bool("decode", false, "")
	digits := flag.Int("digits", 10, "the number of digits for p and q")

	flag.Parse()

	scanner := bufio.NewReader(os.Stdin)

	if *decode {
		data, err := ioutil.ReadFile("sporocilo.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		keys, err := ioutil.ReadFile("kljuci.txt")
		if err != nil {
			fmt.Println(err)
			return
		}

		keysPairs := strings.Split(string(keys), "\n")
		pair := strings.Split(keysPairs[0], ", ")
		e, _ := big.NewInt(0).SetString(pair[0], 10)
		n, _ := big.NewInt(0).SetString(pair[1], 10)

		pair = strings.Split(keysPairs[1], ", ")
		s, _ := big.NewInt(0).SetString(pair[0], 10)

		cipher := cipher.NewCipher(s, e, n)
		decoded := cipher.Decode(string(data))

		fmt.Println(string(decoded))
	} else {
		fmt.Print("Message to encode: ")
		message, err := scanner.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		cipher := cipher.GenerateCipher(*digits)
		cipher.SaveKeys()

		encoded := cipher.Encode(message)

		ioutil.WriteFile("sporocilo.txt", []byte(encoded), 0644)
	}
}
