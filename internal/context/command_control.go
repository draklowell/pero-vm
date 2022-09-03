package context

func (ctx *Context) commandGoto() error {
	offset, err := ctx.readS4()
	if err != nil {
		return err
	}
	ctx.seek(int(offset))
	return nil
}

func (ctx *Context) commandGotoIf() error {
	value, err := ctx.popBoolean()
	if err != nil {
		return err
	}

	if !value {
		return nil
	}

	offset, err := ctx.readS4()
	if err != nil {
		return err
	}
	ctx.seek(int(offset))
	return nil
}
