package context

func (ctx *Context) commandBitwiseAnd() error {
	value1, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value2, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 & value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandBitwiseOr() error {
	value1, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value2, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 | value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandBitwiseXor() error {
	value1, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value2, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 ^ value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandBitwiseNot() error {
	value1, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(^value1)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandBitwiseShiftLeft() error {
	value1, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value2, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 << value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandBitwiseShiftRight() error {
	value1, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value2, err := ctx.popInteger()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 >> value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}
