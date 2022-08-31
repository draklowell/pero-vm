package context

import (
	"lab.draklowell.net/routine-runtime/word"
)

func (ctx *Context) popArray() (*word.Array, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return nil, err
	}

	value, ok := valueWord.(*word.Array)
	if !ok {
		return nil, ErrInvalidWordType
	}

	return value, nil
}

func (ctx *Context) popContainer() (*word.Container, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return nil, err
	}

	value, ok := valueWord.(*word.Container)
	if !ok {
		return nil, ErrInvalidWordType
	}

	return value, nil
}

func (ctx *Context) popInteger() (*word.Integer, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return nil, err
	}
	value, ok := valueWord.(*word.Integer)
	if !ok {
		return nil, ErrInvalidWordType
	}

	return value, nil
}

func (ctx *Context) popFloat() (*word.Float, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return nil, err
	}
	value, ok := valueWord.(*word.Float)
	if !ok {
		return nil, ErrInvalidWordType
	}

	return value, nil
}
func (ctx *Context) popString() (string, error) {
	valueWord, err := ctx.stack.Pop()
	if err != nil {
		return "", err
	}
	value, ok := valueWord.(*word.Bytes)
	if !ok {
		return "", ErrInvalidWordType
	}

	return string(value.GetValue()), nil
}

var (
	TRUE  = int64(1)
	FALSE = int64(0)
)
