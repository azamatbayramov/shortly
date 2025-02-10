package tests

import (
	"github.com/azamatbayramov/shortly/pkg/coder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBaseCoder(t *testing.T) {
	_, err := coder.NewBaseCoder("01", 5)
	assert.NoError(t, err)

	_, err = coder.NewBaseCoder("", 5)
	assert.Error(t, err)

	_, err = coder.NewBaseCoder("00", 5)
	assert.Error(t, err)
}

func TestEncode(t *testing.T) {
	c, _ := coder.NewBaseCoder("0123456789", 5)

	encoded, err := c.Encode(12345)
	assert.NoError(t, err)
	assert.Equal(t, "12345", encoded)

	encoded, err = c.Encode(c.MaxDecodedValue + 1)
	assert.Error(t, err)
}

func TestDecode(t *testing.T) {
	c, _ := coder.NewBaseCoder("0123456789", 5)

	n, err := c.Decode("12345")
	assert.NoError(t, err)
	assert.Equal(t, uint64(12345), n)

	n, err = c.Decode("ABCDE")
	assert.Error(t, err)
}

func TestEncodeDecode(t *testing.T) {
	c, _ := coder.NewBaseCoder("abcdefghijklmnopqrstuvwxyz", 6)

	for i := uint64(0); i < 1000; i++ {
		encoded, err := c.Encode(i)
		assert.NoError(t, err)
		n, err := c.Decode(encoded)
		assert.NoError(t, err)
		assert.Equal(t, i, n)
	}
}
