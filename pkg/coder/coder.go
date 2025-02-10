package coder

type Coder interface {
	Encode(n uint64) (string, error)
	Decode(s string) (uint64, error)
}
