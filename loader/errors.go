package loader

import (
	"errors"
	"fmt"
)

type ErrorUnknownConstantTag struct{ Tag uint8 }

func (e *ErrorUnknownConstantTag) Error() string {
	return fmt.Sprintf("UnknownConstantTag: %d", e.Tag)
}

type ErrorUnknownWordTag struct{ Tag uint8 }

func (e *ErrorUnknownWordTag) Error() string {
	return fmt.Sprintf("UnknownWordTag: %d", e.Tag)
}

type ErrorInvalidMagicNumber struct{ Number []byte }

func (e *ErrorInvalidMagicNumber) Error() string {
	return fmt.Sprintf("InvalidMagicNumber: %X", e.Number)
}

type ErrorUnexpectedEOF struct{ Base error }

func (e *ErrorUnexpectedEOF) Unwrap() error {
	return e.Base
}

func (e *ErrorUnexpectedEOF) Error() string {
	return "UnexpectedEOF"
}

type ErrorIOError struct{ Base error }

func (e *ErrorIOError) Unwrap() error {
	return e.Base
}

func (e *ErrorIOError) Error() string {
	return fmt.Sprintf("IOError[%s]", e.Base.Error())
}

var (
	ErrorInvalidVersion     = errors.New("InvalidVersion")
	ErrorInvalidEntryLength = errors.New("InvalidEntryLength")
)
