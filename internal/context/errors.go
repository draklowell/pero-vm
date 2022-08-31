package context

import (
	"errors"
	"fmt"
)

type ErrorInvalidConstantType struct{ Index uint16 }

func (e *ErrorInvalidConstantType) Error() string {
	return fmt.Sprintf("InvalidConstantType: %d", e.Index)
}

type ErrorConstantNotFound struct{ Index uint16 }

func (e *ErrorConstantNotFound) Error() string {
	return fmt.Sprintf("ConstantNotFound: %d", e.Index)
}

type ErrorTraceBack struct {
	Base   error
	Caller string
	Line   int
}

func (e *ErrorTraceBack) Unwrap() error {
	return e.Base
}

func (e *ErrorTraceBack) Minimal() string {
	baseText := e.Base.Error()
	if trace, ok := e.Base.(*ErrorTraceBack); ok {
		baseText = trace.Minimal()
	}
	return fmt.Sprintf("%s:%d\n%s", e.Caller, e.Line, baseText)
}

func (e *ErrorTraceBack) Error() string {
	return e.Minimal()
}

var (
	ErrorContextFinished    = errors.New("ContextFinishedError")
	ErrorContextBroken      = errors.New("ContextBroken")
	ErrorContextNotFinished = errors.New("ContextNotFinishedError")
	ErrorInvalidWordType    = errors.New("InvalidWordType")
	ErrorInvalidIndex       = errors.New("InvalidIndex")
	ErrorUnexpectedEnd      = errors.New("UnexpectedEnd")
)
