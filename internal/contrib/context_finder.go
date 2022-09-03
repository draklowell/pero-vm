package contrib

import (
	"encoding/binary"

	"lab.draklowell.net/routine-runtime/common/word"
	"lab.draklowell.net/routine-runtime/internal"
	"lab.draklowell.net/routine-runtime/internal/context"
)

type Loader interface {
	GetRoutine(entry string) (binary.ByteOrder, []byte, []context.Constant, map[int]int, string, error)
}

type ContextFinder struct {
	loader Loader
}

func NewContextFinder(loader Loader) *ContextFinder {
	return &ContextFinder{
		loader: loader,
	}
}

func (finder *ContextFinder) Execute(machine *internal.Machine, entry string, arguments []word.Word) ([]word.Word, error) {
	order, bytecode, constants, lineMap, entry, err := finder.loader.GetRoutine(entry)
	if err != nil {
		return nil, err
	}
	if bytecode == nil {
		return nil, nil
	}

	ctx := context.NewContext(machine, order, bytecode, constants, lineMap, entry)

	for index, argument := range arguments {
		ctx.SetVariable(uint8(index), argument)
	}

	err = ctx.Execute(machine.Breaker)
	if err != nil {
		return nil, err
	}

	ret, err := ctx.GetReturn()
	if err != nil {
		return nil, err
	}
	if ret == nil {
		ret = make([]word.Word, 0)
	}
	return ret, nil
}
