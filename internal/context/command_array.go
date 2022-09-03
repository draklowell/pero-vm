package context

import (
	"lab.draklowell.net/routine-runtime/internal/word"
)

func (ctx *Context) commandArrayNew() error {
	size, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.stack.Push(word.NewArray(int(size.GetValue())))
}

func (ctx *Context) commandArrayLength() error {
	array, err := ctx.popArray()
	if err != nil {
		return err
	}

	return ctx.stack.Push(word.NewInteger(int64(array.GetSize())))
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

	return ctx.stack.Push(value)
}

func (ctx *Context) commandArrayGet() error {
	index, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.commandArrayGetStatic(int(index.GetValue()))
}

func (ctx *Context) commandArraySetStatic(index int) error {
	value, err := ctx.stack.Pop()
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

	return ctx.stack.Push(value)
}

func (ctx *Context) commandArraySet() error {
	value, err := ctx.stack.Pop()
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

	return ctx.stack.Push(value)
}
