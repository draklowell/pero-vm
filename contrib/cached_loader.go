package contrib

import (
	"encoding/binary"

	"lab.draklowell.net/routine-runtime/internal/context"
)

type CachedLoader struct {
	Routines map[string]*CachedRoutine
}

func (loader *CachedLoader) GetRoutine(entry string) (binary.ByteOrder, []byte, []context.Constant, map[int]int, string, error) {
	routine := loader.Routines[entry]
	if routine == nil {
		return nil, nil, nil, nil, "", nil
	}

	return routine.Order, routine.Bytecode, routine.Constants, routine.LineMap, routine.Entry, nil
}

type CachedRoutine struct {
	Bytecode  []byte
	Constants []context.Constant
	LineMap   map[int]int
	Entry     string
	Order     binary.ByteOrder
}
