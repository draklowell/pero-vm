package context

func (ctx *Context) commandArrayNew() error {
	size, err := ctx.popInteger()
	if err != nil {
		return err
	}

	array, err := ctx.machine.Heap.NewArray(int(size.GetValue()))
	if err != nil {
		return err
	}
	return ctx.machine.Stack.Push(array)
}

func (ctx *Context) commandArrayLength() error {
	array, err := ctx.popArray()
	if err != nil {
		return err
	}

	integer, err := ctx.machine.Heap.NewInteger(int64(array.GetSize()))
	if err != nil {
		return err
	}
	return ctx.machine.Stack.Push(integer)
}

func (ctx *Context) commandArrayGetStatic(index int) error {
	array, err := ctx.popArray()
	if err != nil {
		return err
	}

	value, err := array.Get(int(index))
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandArrayGet() error {
	index, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.commandArrayGetStatic(int(index.GetValue()))
}

func (ctx *Context) commandArraySetStatic(index int) error {
	value, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}

	array, err := ctx.popArray()
	if err != nil {
		return err
	}

	err = array.Set(int(index), value)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandArraySet() error {
	value, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}

	index, err := ctx.popInteger()
	if err != nil {
		return err
	}

	array, err := ctx.popArray()
	if err != nil {
		return err
	}

	err = array.Set(int(index.GetValue()), value)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}
