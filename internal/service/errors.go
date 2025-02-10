package service

import "errors"

var (
	ErrLinkNotFound           = errors.New("link not found")
	ErrShortLinkIsNotValid    = errors.New("short link is not valid")
	ErrOriginalLinkIsTooLong  = errors.New("original link is too long")
	ErrOriginalLinkIsNotValid = errors.New("original link is not valid")

	ErrEncodeError  = errors.New("encode error")
	ErrStorageError = errors.New("storage error")
)
