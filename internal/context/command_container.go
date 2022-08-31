package context

func (ctx *Context) commandPutFieldStatic(key string) error {
	container, err := ctx.popContainer()
	if err != nil {
		return err
	}

	value, err := ctx.stack.Pop()
	if err != nil {
		return err
	}
	container.Set(key, value)

	return nil
}

func (ctx *Context) commandPutField() error {
	keyIndex, err := ctx.readU2()
	if err != nil {
		return err
	}

	key, err := ctx.getConstantString(keyIndex)
	if err != nil {
		return err
	}
	return ctx.commandPutFieldStatic(key)
}

func (ctx *Context) commandPutFieldDynamic() error {
	key, err := ctx.popString()
	if err != nil {
		return err
	}

	return ctx.commandPutFieldStatic(key)
}
