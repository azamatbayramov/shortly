package coder

import "errors"

var (
	ErrAlphabetLengthLessThanTwo = errors.New("alphabet length less than 2")
	ErrDuplicateCharInAlphabet   = errors.New("duplicate character in alphabet")

	ErrNumberTooLargeToBeEncoded = errors.New("number is too large to be encoded")

	ErrInvalidEncodedStringLength = errors.New("invalid encoded string length")
	ErrInvalidCharInEncodedString = errors.New("invalid character in encoded string")
)
