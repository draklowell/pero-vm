package context

import "lab.draklowell.net/routine-runtime/word"

func (ctx *Context) commandInvokeCollectArguments() ([]word.Word, uint8, error) {
	argumentsCount, err := ctx.readU1()
	if err != nil {
		return nil, 0, err
	}

	retCount, err := ctx.readU1()
	if err != nil {
		return nil, 0, err
	}

	arguments := make([]word.Word, argumentsCount)
	for i := uint8(0); i < argumentsCount; i++ {
		arguments[i], err = ctx.stack.Pop()
		if err != nil {
			return nil, 0, err
		}
	}

	return arguments, retCount, nil
}

func (ctx *Context) commandInvokeCollectReturn(ret []word.Word, retCount uint8) error {
	for i := uint8(0); i < retCount; i++ {
		var value word.Word
		if len(ret) > int(i) {
			value = ret[i]
		}
		if value == nil {
			value = word.None
		}
		err := ctx.stack.Push(value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ctx *Context) commandInvokeDynamic() error {
	arguments, retCount, err := ctx.commandInvokeCollectArguments()
	if err != nil {
		return err
	}

	entry, err := ctx.popString()
	if err != nil {
		return err
	}

	ret, err := ctx.machine.Execute(ctx.GetCaller(), entry, arguments)
	if err != nil {
		return err
	}

	return ctx.commandInvokeCollectReturn(ret, retCount)
}

func (ctx *Context) commandInvoke() error {
	entryIndex, err := ctx.readU2()
	if err != nil {
		return err
	}

	arguments, retCount, err := ctx.commandInvokeCollectArguments()
	if err != nil {
		return err
	}

	entry, err := ctx.getConstantString(entryIndex)
	if err != nil {
		return err
	}

	ret, err := ctx.machine.Execute(ctx.GetCaller(), entry, arguments)
	if err != nil {
		return err
	}

	return ctx.commandInvokeCollectReturn(ret, retCount)
}
