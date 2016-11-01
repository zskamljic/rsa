package util

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtendedEuclid(t *testing.T) {
	assert := assert.New(t)

	// ARRANGE
	a := big.NewInt(5)
	b := big.NewInt(10)

	// ACT
	d, x, y := ExtendedEuclid(a, b)

	// ASSERT
	assert.Equal(big.NewInt(5), d)
	assert.Equal(big.NewInt(1), x)
	assert.Equal(big.NewInt(0), y)
}

func TestModLinEquation(t *testing.T) {
	assert := require.New(t)

	// ARRANGE
	a := big.NewInt(3)
	n := big.NewInt(2)

	// ACT
	x := ModLinEquation(a, n)

	// ASSERT
	assert.NotNil(x)
	assert.Equal(big.NewInt(1), x)
}
