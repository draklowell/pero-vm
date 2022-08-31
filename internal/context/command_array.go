package context

import (
	"lab.draklowell.net/routine-runtime/word"
)

func absoluteIndex(x int64, length int) int64 {
	if x < 0 {
		x = int64(length) - x - 1
	}
	if x < 0 {
		return int64(length) // Remake pls :D
	}
	return x
}

// -> array[array]
func (ctx *Context) commandArrayNew() error {
	size, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.stack.Push(word.NewArray(int(size.GetValue())))
}

// array[array] -> length[int64]
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

// array[array], index[int64] -> element[any]
func (ctx *Context) commandArrayGet() error {
	index, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.commandArrayGetStatic(int(index.GetValue()))
}
