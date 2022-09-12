package context

import (
	"lab.draklowell.net/pero-core/common/word"
)

func (ctx *Context) popArray() (*word.Array, error) {
	valueWord, err := ctx.machine.Stack.Pop()
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
	valueWord, err := ctx.machine.Stack.Pop()
	if err != nil {
		return nil, err
	}

	value, ok := valueWord.(*word.Container)
	if !ok {
		return nil, ErrInvalidWordType
	}

	return value, nil
}

func (ctx *Context) popInteger() (int64, error) {
	valueWord, err := ctx.machine.Stack.Pop()
	if err != nil {
		return 0, err
	}
	value, ok := valueWord.(*word.Integer)
	if !ok {
		return 0, ErrInvalidWordType
	}

	return value.GetValue(), nil
}

func (ctx *Context) popString() (string, error) {
	valueWord, err := ctx.machine.Stack.Pop()
	if err != nil {
		return "", err
	}
	value, ok := valueWord.(*word.Bytes)
	if !ok {
		return "", ErrInvalidWordType
	}

	return string(value.GetValue()), nil
}

func (ctx *Context) popBoolean() (bool, error) {
	valueWord, err := ctx.machine.Stack.Pop()
	if err != nil {
		return false, err
	}
	value, ok := valueWord.(*word.Boolean)
	if !ok {
		return false, ErrInvalidWordType
	}

	return value.GetValue(), nil
}

func (ctx *Context) popIntegerSoft() (int64, error) {
	valueWord, err := ctx.machine.Stack.Pop()
	if err != nil {
		return 0, err
	}

	switch value := valueWord.(type) {
	case *word.Integer:
		return value.GetValue(), nil
	case *word.Float:
		return int64(value.GetValue()), nil
	}

	return 0, ErrInvalidWordType
}

func (ctx *Context) popFloatSoft() (float64, error) {
	valueWord, err := ctx.machine.Stack.Pop()
	if err != nil {
		return 0, err
	}

	switch value := valueWord.(type) {
	case *word.Integer:
		return float64(value.GetValue()), nil
	case *word.Float:
		return value.GetValue(), nil
	}

	return 0, ErrInvalidWordType
}

var (
	TRUE  = int64(1)
	FALSE = int64(0)
)
