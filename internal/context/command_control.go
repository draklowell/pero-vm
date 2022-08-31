package context

func (ctx *Context) commandGoto() error {
	offset, err := ctx.readS4()
	if err != nil {
		return err
	}
	ctx.seek(int(offset))
	return nil
}
