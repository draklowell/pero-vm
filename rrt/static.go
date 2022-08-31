package rrt

import (
	"io"

	"lab.draklowell.net/routine-runtime/contrib"
	"lab.draklowell.net/routine-runtime/loader"
)

func (vm *VirtualMachine) AddStaticRoutine(stream io.Reader) (string, error) {
	bytecode, constants, lineMap, entry, order, err := loader.LoadRoutine(stream)
	if err != nil {
		return "", err
	}

	vm.staticLoader.Routines[entry] = &contrib.CachedRoutine{
		Bytecode:  bytecode,
		Constants: constants,
		LineMap:   lineMap,
		Entry:     entry,
		Order:     order,
	}

	return entry, nil
}
