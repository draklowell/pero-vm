package context

import "lab.draklowell.net/pero-core/common/word"

func (ctx *Context) commandPushConstant() error {
	index, err := ctx.readU2()
	if err != nil {
		return err
	}

	word, err := ctx.getConstantWord(index)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(word)
}

func (ctx *Context) commandPushNone() error {
	return ctx.machine.Stack.Push(nil)
}

func (ctx *Context) commandPushIntegerStatic(value int64) error {
	return ctx.machine.Stack.Push(word.NewInteger(value))
}

func (ctx *Context) commandPushFloatStatic(value float64) error {
	return ctx.machine.Stack.Push(word.NewFloat(value))
}

func (ctx *Context) commandPushVariable() error {
	index, err := ctx.readU1()
	if err != nil {
		return err
	}

	return ctx.commandPushVariableStatic(index)
}

func (ctx *Context) commandPushVariableStatic(index uint8) error {
	return ctx.machine.Stack.Push(ctx.GetVariable(index))
}

func (ctx *Context) commandPop() error {
	return ctx.machine.Stack.Push(nil)
}

func (ctx *Context) commandPopVariable() error {
	index, err := ctx.readU1()
	if err != nil {
		return err
	}

	return ctx.commandPopVariableStatic(index)
}

func (ctx *Context) commandPopVariableStatic(index uint8) error {
	word, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}
	ctx.SetVariable(index, word)
	return nil
}
