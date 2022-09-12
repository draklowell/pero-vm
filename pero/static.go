package pero

import (
	"bytes"

	"lab.draklowell.net/pero-core/internal/contrib"
	"lab.draklowell.net/pero-core/loader"
)

func (vm *VirtualMachine) AddStaticRoutine(data []byte) (string, error) {
	bytecode, constants, lineMap, entry, order, err := loader.LoadRoutine(bytes.NewReader(data))
	if err != nil {
		return "", err
	}

	vm.staticLoader.Routines[entry] = &contrib.StaticRoutine{
		Bytecode:  bytecode,
		Constants: constants,
		LineMap:   lineMap,
		Entry:     entry,
		Order:     order,
	}

	return entry, nil
}
