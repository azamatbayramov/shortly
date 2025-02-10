package coder

import (
	"math"
	"strings"
)

type BaseCoder struct {
	alphabet        string
	base            int
	length          int
	charToIndex     map[rune]int
	MaxDecodedValue uint64
}

var _ Coder = (*BaseCoder)(nil)

func NewBaseCoder(alphabet string, length int) (*BaseCoder, error) {
	if len(alphabet) < 2 {
		return nil, ErrAlphabetLengthLessThanTwo
	}

	charToIndex := make(map[rune]int)

	for i, char := range alphabet {
		if _, exists := charToIndex[char]; exists {
			return nil, ErrDuplicateCharInAlphabet
		}
		charToIndex[char] = i
	}

	maxDecodedValue := uint64(math.Pow(float64(len(alphabet)), float64(length))) - 1

	return &BaseCoder{
		alphabet:        alphabet,
		base:            len(alphabet),
		length:          length,
		charToIndex:     charToIndex,
		MaxDecodedValue: maxDecodedValue,
	}, nil
}

func (coder BaseCoder) Encode(n uint64) (string, error) {
	var encoded strings.Builder

	if n > coder.MaxDecodedValue {
		return "", ErrNumberTooLargeToBeEncoded
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
		return 0, ErrInvalidEncodedStringLength
	}

	for _, char := range s {
		index, exists := coder.charToIndex[char]

		if !exists {
			return 0, ErrInvalidCharInEncodedString
		}

		n = n*uint64(coder.base) + uint64(index)
	}

	return n, nil
}
