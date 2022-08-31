package word

import "errors"

var (
	ErrInvalidIndex = errors.New("invalid array index")
	ErrInvalidKey   = errors.New("invalid container key")
	ErrNilPointer   = errors.New("unexpected nil pointer")
)
