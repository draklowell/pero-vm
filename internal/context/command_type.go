package context

import "lab.draklowell.net/routine-runtime/word"

func (ctx *Context) commandCheckTypeStatic(typ int) error {
	value, err := ctx.stack.Pop()
	if err != nil {
		return err
	}

	if word.GetType(value) == typ {
		ctx.stack.Push(TRUE)
	} else {
		ctx.stack.Push(FALSE)
	}

	return nil
}

func (ctx *Context) commandFloatToInteger() error {
	value, err := ctx.popFloat()
	if err != nil {
		return err
	}

	return ctx.stack.Push(int64(value))
}

func (ctx *Context) commandIntegerToFloat() error {
	value, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.stack.Push(float64(value))
}
