package internal

import (
	"errors"
	"fmt"
)

type ErrEntryNotFound struct{ Entry string }

func (e *ErrEntryNotFound) Error() string {
	return fmt.Sprintf("entry \"%s\" not found", e.Entry)
}

var (
	ErrStackTooLarge = errors.New("stack too large")
	ErrStackEmpty    = errors.New("stack empty")
	ErrNilPointer    = errors.New("unexpected nil pointer")
)
