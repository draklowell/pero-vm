package rrt

import (
	"errors"
	"fmt"

	"lab.draklowell.net/routine-runtime/internal/context"
)

type Unwrapable interface {
	Error() string
	Unwrap() error
}

type ErrExecution struct {
	Base Unwrapable
}

func (e *ErrExecution) Unwrap() error {
	return e.Base
}

func (e *ErrExecution) InvokeBackTrace() ([]BackTraceElement, error) {
	trace := make([]BackTraceElement, 0)

	var current Unwrapable = e
	for {
		unwraped := current.Unwrap()
		if err, ok := unwraped.(*context.ErrTraceBack); ok {
			trace = append(trace, BackTraceElement{
				Entry: err.Caller,
				Line:  err.Line,
			})
		}
		if err, ok := unwraped.(Unwrapable); ok {
			current = err
		} else {
			break
		}
	}

	return trace, current.Unwrap()
}

func (e *ErrExecution) Error() string {
	trace, base := e.InvokeBackTrace()

	result := base.Error()

	for _, element := range trace {
		result += fmt.Sprintf("\n    In %s at line %d", element.Entry, element.Line)
	}

	return result
}

type BackTraceElement struct {
	Entry string
	Line  int
}

var (
	ErrInvalidEntry        = errors.New("invalid entry")
	ErrLoaderLimitExceeded = errors.New("loader limit exceeded")
)
