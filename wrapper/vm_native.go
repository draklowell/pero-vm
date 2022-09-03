package main

import (
	"lab.draklowell.net/routine-runtime/common/word"
	"lab.draklowell.net/routine-runtime/rrt"
)

//#include <types.h>
import "C"

type nativeRoutineWrapper struct {
	base C.rrtNativeRoutine
}

func (nrw *nativeRoutineWrapper) Run(machine *rrt.VirtualMachine, arguments []word.Word) ([]word.Word, error) {
	vmPtr := vms.Find(
		func(value *VM) bool {
			return value != nil && value.machine == machine
		},
	)
	if vmPtr == NullPointer {
		return nil, nil
	}

	vm, _ := vms.Get(vmPtr)

	ptrArguments := make([]C.int, len(arguments))
	for i, argument := range arguments {
		argumentPtr, err := vm.words.Add(argument)
		if err != nil {
			return nil, err
		}
		ptrArguments[i] = C.int(argumentPtr)
	}
	var ptrArgumentSize C.int = 0
	ptrArgumentsPtr := CArray(ptrArguments, &ptrArgumentSize)

	var cRetSize C.int = 0
	cRet, err := C.rrtNativeRoutineBridge(
		nrw.base,
		C.int(vmPtr),
		ptrArgumentsPtr,
		ptrArgumentSize,
		&cRetSize,
	)
	if err != nil {
		return nil, err
	}

	rets := make([]word.Word, 0)
	if cRet != nil && cRetSize != 0 {
		cRets := GoArray(cRet, cRetSize)
		rets = make([]word.Word, len(cRets))
		for i, ret := range cRets {
			value, err := vm.words.Get(Pointer(ret))
			if err != nil {
				return nil, err
			}
			rets[i] = value
		}
	}

	for _, argument := range ptrArguments {
		vm.words.Remove(Pointer(argument))
	}

	return rets, nil
}

//export rrtVMAddNativeRoutine
func rrtVMAddNativeRoutine(vmPtr Pointer, entry *C.char, routine C.rrtNativeRoutine) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	wrapper := &nativeRoutineWrapper{base: routine}
	vm.machine.AddNativeRoutine(C.GoString(entry), wrapper.Run)
	return 0
}
