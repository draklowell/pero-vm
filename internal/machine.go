package internal

import (
	"lab.draklowell.net/routine-runtime/word"
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

	StackSize uint16
	Breaker   BreakCallback
}

func NewMachine(finder ModuleFinder, traceSize uint16, stackSize uint16) *Machine {
	return &Machine{
		StackSize: stackSize,
		finder:    finder,
		Breaker:   EmptyBreaker,
	}
}

func (machine *Machine) Execute(caller string, entry string, arguments []word.Word) ([]word.Word, error) {
	ret, err := machine.finder.Execute(machine, entry, arguments)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		return nil, &ErrEntryNotFound{Entry: entry}
	}

	return ret, nil
}
