package context

func (ctx *Context) commandLogicalAnd() error {
	value1, err := ctx.popBoolean()
	if err != nil {
		return err
	}

	value2, err := ctx.popBoolean()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 && value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandLogicalOr() error {
	value1, err := ctx.popBoolean()
	if err != nil {
		return err
	}

	value2, err := ctx.popBoolean()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 || value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandLogicalNot() error {
	value1, err := ctx.popBoolean()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(!value1)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}
