package internal

import (
	"strings"

	"lab.draklowell.net/routine-runtime/internal/word"
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
}

func NewMachine(finder ModuleFinder, stackSize uint, heapSize uint) *Machine {
	return &Machine{
		finder: finder,

		Breaker: EmptyBreaker,
		Stack:   NewStack(stackSize),
	}
}

func (machine *Machine) Execute(caller string, entry string, arguments []word.Word) ([]word.Word, error) {
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
