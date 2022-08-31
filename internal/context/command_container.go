package context

func (ctx *Context) commandPutFieldStatic(name string) error {
	container, err := ctx.popContainer()
	if err != nil {
		return err
	}

	value, err := ctx.stack.Pop()
	if err != nil {
		return err
	}
	container[name] = value

	return nil
}

func (ctx *Context) commandPutField() error {
	nameIndex, err := ctx.readU2()
	if err != nil {
		return err
	}

	name, err := ctx.getConstantString(nameIndex)
	if err != nil {
		return err
	}
	return ctx.commandPutFieldStatic(name)
}

func (ctx *Context) commandPutFieldDynamic() error {
	name, err := ctx.popString()
	if err != nil {
		return err
	}

	return ctx.commandPutFieldStatic(name)
}
