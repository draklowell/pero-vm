package contrib

import (
	"encoding/binary"

	"lab.draklowell.net/pero-core/internal/context"
)

type StaticLoader struct {
	Routines map[string]*StaticRoutine
}

func (loader *StaticLoader) GetRoutine(entry string) (binary.ByteOrder, []byte, []context.Constant, map[int]int, string, error) {
	routine := loader.Routines[entry]
	if routine == nil {
		return nil, nil, nil, nil, "", nil
	}

	return routine.Order, routine.Bytecode, routine.Constants, routine.LineMap, routine.Entry, nil
}

type StaticRoutine struct {
	Bytecode  []byte
	Constants []context.Constant
	LineMap   map[int]int
	Entry     string
	Order     binary.ByteOrder
}
