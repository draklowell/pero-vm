package utils

import "errors"

var (
	ErrStackTooLarge = errors.New("stack too large")
	ErrStackEmpty    = errors.New("stack empty")
)
