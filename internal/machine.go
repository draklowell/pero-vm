package internal

import (
	"runtime"
	"strings"

	"lab.draklowell.net/routine-runtime/common/word"
)

const (
	GCFrequent = 0
	GCRare     = 1
)

type BreakCallback func() bool

func EmptyBreaker() bool {
	return false
}

type ModuleFinder interface {
	Execute(*Machine, string, []word.Word) ([]word.Word, error)
}

type Machine struct {
	finder ModuleFinder

	Breaker BreakCallback
	Stack   *Stack
	Heap    *Heap

	gcMode int
}

func NewMachine(finder ModuleFinder, stackSize uint, heapSize uint, gcMode int) *Machine {
	return &Machine{
		finder: finder,

		Breaker: EmptyBreaker,
		Stack:   NewStack(stackSize),
		Heap:    NewHeap(heapSize, gcMode),

		gcMode: gcMode,
	}
}

func (machine *Machine) Execute(caller string, entry string, arguments []word.Word) ([]word.Word, error) {
	if machine.gcMode == GCFrequent {
		runtime.GC()
	}

	if strings.ContainsRune(entry, 0) {
		return nil, &ErrEntryNotFound{Entry: entry}
	}

	ret, err := machine.finder.Execute(machine, entry, arguments)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return nil, &ErrEntryNotFound{Entry: entry}
	}

	return ret, nil
}
