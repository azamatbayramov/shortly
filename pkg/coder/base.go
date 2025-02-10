package coder

import (
	"errors"
	"math"
	"strings"
)

type BaseCoder struct {
	alphabet        string
	base            int
	length          int
	charToIndex     map[rune]int
	maxDecodedValue uint64
}

var _ Coder = (*BaseCoder)(nil)

func NewBaseCoder(alphabet string, length int) (*BaseCoder, error) {
	if len(alphabet) < 2 {
		return nil, errors.New("alphabet length must be at least 2")
	}

	charToIndex := make(map[rune]int)

	for i, char := range alphabet {
		if _, exists := charToIndex[char]; exists {
			return nil, errors.New("duplicate character in alphabet")
		}
		charToIndex[char] = i
	}

	maxDecodedValue := uint64(math.Pow(float64(len(alphabet)), float64(length))) - 1

	return &BaseCoder{
		alphabet:        alphabet,
		base:            len(alphabet),
		length:          length,
		charToIndex:     charToIndex,
		maxDecodedValue: maxDecodedValue,
	}, nil
}

func (coder BaseCoder) Encode(n uint64) (string, error) {
	var encoded strings.Builder

	if n > coder.maxDecodedValue {
		return "", errors.New("number is too large to be encoded")
	}

	for n > 0 {
		rem := n % uint64(coder.base)
		n /= uint64(coder.base)
		encoded.WriteByte(coder.alphabet[rem])
	}

	if encoded.Len() < coder.length {
		encoded.WriteString(strings.Repeat(string(coder.alphabet[0]), coder.length-encoded.Len()))
	}

	runes := []rune(encoded.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes), nil
}

func (coder BaseCoder) Decode(s string) (uint64, error) {
	var n uint64

	if len(s) != coder.length {
		return 0, errors.New("invalid encoded string length")
	}

	for _, char := range s {
		index, exists := coder.charToIndex[char]

		if !exists {
			return 0, errors.New("invalid character in encoded string")
		}

		n = n*uint64(coder.base) + uint64(index)
	}

	return n, nil
}
