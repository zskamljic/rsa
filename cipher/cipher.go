package cipher

import (
	"bufio"
	"fmt"
	"math/big"
	"strings"

	"os"

	"github.com/zskamljic/rsa/util"
)

// Cipher is used to encipher or decipher RSA
type Cipher struct {
	d *big.Int
	e *big.Int
	n *big.Int
}

// NewCipher returns a new Cipher
func NewCipher(d, e, n *big.Int) *Cipher {
	return &Cipher{d, e, n}
}

// GenerateCipher generates a cipher with p and q
// containing @digits digits.
func GenerateCipher(digits int) *Cipher {
	var p, q *big.Int

	// Step 1: select two similar primes p and q
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

	// Step 2: φ(n) = (p-1)(q-1)
	euler := big.NewInt(0)
	euler.Mul(p, q)

	// Step 3: Select such e that 1<e<φ(n) and gcd(e,φ(n)) == 1
	var e *big.Int
	gcd := big.NewInt(0)
	for {
		e = generatePrimeMiller(digits - 1)

		if gcd.GCD(nil, nil, e, euler).Cmp(big.NewInt(1)) == 0 {
			break
		}
	}

	//Step 4: calculate d such that ed≡1(mod φ(n))
	d := util.ModLinEquation(e, euler)

	return NewCipher(d, e, n)
}

// SaveKeys saves the keys to "kljuci.txt"
func (c *Cipher) SaveKeys() {
	out, err := os.Create("kljuci.txt")
	if err != nil {
		fmt.Println("Unable to open \"kljuci.txt\":", err)
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	defer writer.Flush()

	writer.WriteString(fmt.Sprintf("Public: %v, %v\n", c.e, c.n))
	writer.WriteString(fmt.Sprintf("Private: %v, %v\n", c.d, c.n))
}

// Encode encodes the message to []byte
func (c *Cipher) Encode(message string) string {
	messageBytes := []byte(message)

	msgLen := int64(len(message))

	out := fmt.Sprintf("%v", util.ModExp(big.NewInt(msgLen), c.e, c.n))

	for _, v := range messageBytes {
		m := big.NewInt(int64(v))
		encoded := util.ModExp(m, c.e, c.n)

		out += fmt.Sprintf(" %v", encoded)
		//bytes := encoded.Bytes()
		//out = append(out, byte(len(bytes)))
		//out = append(out, bytes...)
	}
	//m := big.NewInt(0).SetBytes(messageBytes)
	//fmt.Println(m)
	//encoded := util.ModExp(m, c.e, c.n)
	//decoded := util.ModExp(encoded, c.d, c.n)
	//fmt.Println(decoded)

	return out
}

// Decode decodes the message from data
func (c *Cipher) Decode(data string) []byte {

	numbers := strings.Fields(data)
	msgLen, _ := big.NewInt(0).SetString(numbers[0], 10)
	msgLen = util.ModExp(msgLen, c.d, c.n)
	numbers = numbers[1:]

	var message []byte
	for len(numbers) > 0 {
		number, _ := big.NewInt(0).SetString(numbers[0], 10)
		numbers = numbers[1:]

		number = util.ModExp(number, c.d, c.n)
		message = append(message, number.Bytes()...)
	}

	if msgLen.Cmp(big.NewInt(int64(len(message)))) != 0 {
		return nil
	}

	return message
}
