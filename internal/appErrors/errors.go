package appErrors

import "errors"

var ShortLinkIsNotValid = errors.New("short link is not valid")

var OriginalLinkIsTooLong = errors.New("original link is too long")
var OriginalLinkIsNotValid = errors.New("original link is not valid")

var LinkNotFound = errors.New("link not found")

var EncodeError = errors.New("encode error")
var StorageError = errors.New("storage error")
