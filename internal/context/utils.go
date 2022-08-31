package context

import (
	"lab.draklowell.net/routine-runtime/word"
)

func (ctx *Context) popArray() ([]word.Word, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return nil, err
	}

	value, ok := valueWord.([]word.Word)
	if !ok {
		return nil, ErrorInvalidWordType
	}

	return value, nil
}

func (ctx *Context) popContainer() (map[string]word.Word, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return nil, err
	}

	value, ok := valueWord.(map[string]word.Word)
	if !ok {
		return nil, ErrorInvalidWordType
	}

	return value, nil
}

func (ctx *Context) popInteger() (int64, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return 0, err
	}
	value, ok := valueWord.(int64)
	if !ok {
		return 0, ErrorInvalidWordType
	}

	return value, nil
}

func (ctx *Context) popFloat() (float64, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return 0, err
	}
	value, ok := valueWord.(float64)
	if !ok {
		return 0, ErrorInvalidWordType
	}

	return value, nil
}
func (ctx *Context) popString() (string, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return "", err
	}
	value, ok := valueWord.([]uint8)
	if !ok {
		return "", ErrorInvalidWordType
	}

	return string(value), nil
}

var (
	TRUE  = int64(1)
	FALSE = int64(0)
)
