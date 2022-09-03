package context

func (ctx *Context) commandCheckTypeStatic(typ int) error {
	value, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}

	boolean, err := ctx.machine.Heap.NewBoolean(value.GetType() == typ)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(boolean)
}

func (ctx *Context) commandFloatToInteger() error {
	value, err := ctx.popFloat()
	if err != nil {
		return err
	}

	integer, err := ctx.machine.Heap.NewInteger(int64(value.GetValue()))
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(integer)
}

func (ctx *Context) commandIntegerToFloat() error {
	value, err := ctx.popInteger()
	if err != nil {
		return err
	}

	float, err := ctx.machine.Heap.NewFloat(float64(value.GetValue()))
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(float)
}
