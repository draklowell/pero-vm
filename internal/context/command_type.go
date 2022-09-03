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

func (ctx *Context) commandCheckNotTypeStatic(typ int) error {
	value, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}

	boolean, err := ctx.machine.Heap.NewBoolean(value.GetType() != typ)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(boolean)
}

func (ctx *Context) commandToInteger() error {
	value, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	integer, err := ctx.machine.Heap.NewInteger(int64(value))
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(integer)
}

func (ctx *Context) commandToFloat() error {
	value, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	float, err := ctx.machine.Heap.NewFloat(float64(value))
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(float)
}
