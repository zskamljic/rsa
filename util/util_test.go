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

func TestModExpPreservesValues(t *testing.T) {
	assert := assert.New(t)

	// ARRANGE
	aOrig := big.NewInt(8)
	bOrig := big.NewInt(868)
	mOrig := big.NewInt(100)

	a := big.NewInt(0).Set(aOrig)
	b := big.NewInt(0).Set(bOrig)
	m := big.NewInt(0).Set(mOrig)
	res := big.NewInt(16)

	// ACT
	x := ModExp(a, b, m)

	// ASSERT
	assert.Equal(aOrig, a)
	assert.Equal(bOrig, b)
	assert.Equal(mOrig, m)
	assert.Equal(res, x)
}
