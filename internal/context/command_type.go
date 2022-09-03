package context

import "lab.draklowell.net/routine-runtime/internal/word"

func (ctx *Context) commandCheckTypeStatic(typ int) error {
	value, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}

	ctx.machine.Stack.Push(word.NewBoolean(value.GetType() == typ))

	return nil
}

func (ctx *Context) commandFloatToInteger() error {
	value, err := ctx.popFloat()
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(word.NewInteger(int64(value.GetValue())))
}

func (ctx *Context) commandIntegerToFloat() error {
	value, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(word.NewFloat(float64(value.GetValue())))
}
