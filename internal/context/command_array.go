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

	return ctx.stack.Push(make([]word.Word, size))
}

// array[array] -> length[int64]
func (ctx *Context) commandArrayLength() error {
	array, err := ctx.popArray()
	if err != nil {
		return err
	}

	return ctx.stack.Push(int64(len(array)))
}

func (ctx *Context) commandArrayGetStatic(index int64) error {
	array, err := ctx.popArray()
	if err != nil {
		return err
	}

	index = absoluteIndex(index, len(array))

	if index > int64(len(array)) {
		return ErrorInvalidIndex
	}
	return ctx.stack.Push(array[index])
}

// array[array], index[int64] -> element[any]
func (ctx *Context) commandArrayGet() error {
	index, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.commandArrayGetStatic(index)
}
func (ctx *Context) commandArraySetStatic(index int64) error {
	array, err := ctx.popArray()
	if err != nil {
		return err
	}

	index = absoluteIndex(index, len(array))

	if index > int64(len(array)) {
		return ErrorInvalidIndex
	}
	return ctx.stack.Push(array[index])
}

// array[array], index[int64] -> element[any]
func (ctx *Context) commandArraySet() error {
	index, err := ctx.popInteger()
	if err != nil {
		return err
	}

	return ctx.commandArraySetStatic(index)
}
