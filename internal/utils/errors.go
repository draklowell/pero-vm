package utils

import "errors"

var (
	ErrorStackTooLarge = errors.New("StackTooLarge")
	ErrorStackEmpty    = errors.New("StackEmpty")
)
