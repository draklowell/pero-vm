package context

import "lab.draklowell.net/routine-runtime/word"

func (ctx *Context) commandCheckTypeStatic(typ int) error {
	value, err := ctx.stack.Pop()
	if err != nil {
		return err
	}

	ctx.stack.Push(word.NewBoolean(value.GetType() == typ))

	return nil
}

func (ctx *Context) commandFloatToInteger() error {
	value, err := ctx.popFloat()
	if err != nil {
		return err
	}

	return ctx.stack.Push(word.NewInteger(int64(value.GetValue())))
}

func (ctx *Context) commandIntegerToFloat() error {
	value, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.stack.Push(word.NewFloat(float64(value.GetValue())))
}
