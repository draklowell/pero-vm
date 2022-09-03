package context

func (ctx *Context) commandCompareIGreater() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 > value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareIGreaterEquals() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 >= value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareILower() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 < value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareILowerEquals() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 <= value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareIEquals() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 == value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareINotEquals() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 != value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareFGreater() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 > value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareFGreaterEquals() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 >= value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareFLower() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 < value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareFLowerEquals() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 <= value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareFEquals() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 == value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandCompareFNotEquals() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewBoolean(value1 != value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}
