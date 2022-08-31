package loader

import (
	"errors"
	"fmt"
)

type ErrUnknownConstantTag struct{ Tag uint8 }

func (e *ErrUnknownConstantTag) Error() string {
	return fmt.Sprintf("unknown constant tag @%d", e.Tag)
}

type ErrUnknownWordTag struct{ Tag uint8 }

func (e *ErrUnknownWordTag) Error() string {
	return fmt.Sprintf("unknown word tag @%d", e.Tag)
}

type ErrInvalidMagicNumber struct{ Number []byte }

func (e *ErrInvalidMagicNumber) Error() string {
	return fmt.Sprintf("invalid magic number %X", e.Number)
}

type ErrUnexpectedEOF struct{ Base error }

func (e *ErrUnexpectedEOF) Unwrap() error {
	return e.Base
}

func (e *ErrUnexpectedEOF) Error() string {
	return "unexpected EOF"
}

type ErrIOError struct{ Base error }

func (e *ErrIOError) Unwrap() error {
	return e.Base
}

func (e *ErrIOError) Error() string {
	return fmt.Sprintf("IOError %s", e.Base.Error())
}

var (
	ErrInvalidVersion     = errors.New("invalid version")
	ErrInvalidEntryLength = errors.New("invalid entry length")
)
