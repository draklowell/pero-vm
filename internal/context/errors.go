package context

import (
	"errors"
	"fmt"
)

type ErrInvalidConstantType struct{ Index uint16 }

func (e *ErrInvalidConstantType) Error() string {
	return fmt.Sprintf("constant type @%d invalid", e.Index)
}

type ErrConstantNotFound struct{ Index uint16 }

func (e *ErrConstantNotFound) Error() string {
	return fmt.Sprintf("constant #%d not found", e.Index)
}

type ErrTraceBack struct {
	Base   error
	Caller string
	Line   int
}

func (e *ErrTraceBack) Unwrap() error {
	return e.Base
}

func (e *ErrTraceBack) Minimal() string {
	baseText := e.Base.Error()
	if trace, ok := e.Base.(*ErrTraceBack); ok {
		baseText = trace.Minimal()
	}
	return fmt.Sprintf("%s:%d\n%s", e.Caller, e.Line, baseText)
}

func (e *ErrTraceBack) Error() string {
	return e.Minimal()
}

var (
	ErrContextFinished    = errors.New("context finished")
	ErrContextBroken      = errors.New("context broken")
	ErrContextNotFinished = errors.New("context not finished yet")
	ErrInvalidWordType    = errors.New("invalid word type")
	ErrUnexpectedEnd      = errors.New("unexpected bytecode end")
	ErrStackTooLarge      = errors.New("stack too large")
	ErrStackEmpty         = errors.New("stack empty")
)
