package appErrors

import "errors"

var (
	ShortLinkIsNotValid    = errors.New("short link is not valid")
	OriginalLinkIsTooLong  = errors.New("original link is too long")
	OriginalLinkIsNotValid = errors.New("original link is not valid")
	LinkNotFound           = errors.New("link not found")
	EncodeError            = errors.New("encode error")
	StorageError           = errors.New("storage error")
)
