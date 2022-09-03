package context

import (
	"math"

	"lab.draklowell.net/routine-runtime/common/word"
)

func (ctx *Context) commandMathIAdd() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 + value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathISub() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 - value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathIMul() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 * value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathIDiv() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 / value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathIRem() error {
	value1, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popIntegerSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewInteger(value1 % value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathFAdd() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewFloat(value1 + value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathFSub() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewFloat(value1 - value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathFMul() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewFloat(value1 * value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathFDiv() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewFloat(value1 / value2)
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathFRem() error {
	value1, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value2, err := ctx.popFloatSoft()
	if err != nil {
		return err
	}

	value, err := ctx.machine.Heap.NewFloat(math.Mod(value1, value2))
	if err != nil {
		return err
	}

	return ctx.machine.Stack.Push(value)
}

func (ctx *Context) commandMathNeg() error {
	valueWord, err := ctx.machine.Stack.Pop()
	if err != nil {
		return err
	}

	switch value := valueWord.(type) {
	case *word.Integer:
		result, err := ctx.machine.Heap.NewInteger(-value.GetValue())
		if err != nil {
			return err
		}

		return ctx.machine.Stack.Push(result)
	case *word.Float:
		result, err := ctx.machine.Heap.NewFloat(-value.GetValue())
		if err != nil {
			return err
		}

		return ctx.machine.Stack.Push(result)
	}

	return ErrInvalidWordType
}
