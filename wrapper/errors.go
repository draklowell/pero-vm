package main

import "errors"

import "C"

var errorMap = map[int64]error{}

func throw(err error) {
	errorMap[getThreadId()] = err
}

//export rrtGetError
func rrtGetError() *C.char {
	return C.CString(errorMap[getThreadId()].Error())
}

var (
	ErrInvalidWordType         = errors.New("invalid word type")
	ErrInvalidPointer          = errors.New("invalid pointer")
	ErrPointerSetLimitExceeded = errors.New("pointer set limit exceeded")
	ErrInvalidPointerSetSize   = errors.New("invalid pointer set size")
)
