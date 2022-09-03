package main

import (
	"lab.draklowell.net/routine-runtime/internal/word"
	"lab.draklowell.net/routine-runtime/rrt"
)

//#include <types.h>
import "C"

type VM struct {
	machine *rrt.VirtualMachine
	words   PointerSet[word.Word]
}

var vms, _ = NewPointerSet[*VM](128)

//export rrtVMNew
func rrtVMNew(name *C.char, wordPointerSetSize C.int, stackSize C.int, heapSize C.int) Pointer {
	wordsCache, err := NewPointerSet[word.Word](int(wordPointerSetSize))
	if err != nil {
		throw(err)
		return NullPointer
	}

	ptr, err := vms.Add(&VM{
		machine: rrt.NewVirtualMachine(
			C.GoString(name),
			uint(stackSize),
			uint(heapSize),
		),
		words: wordsCache,
	})

	if err != nil {
		throw(err)
		return NullPointer
	}
	return ptr
}

//export rrtVMClose
func rrtVMClose(vmPtr Pointer) C.int {
	err := vms.Remove(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}
	return 0
}

//export rrtVMRemoveRoutine
func rrtVMRemoveRoutine(vmPtr Pointer, entry *C.char) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	vm.machine.RemoveRoutine(C.GoString(entry))
	return 0
}

//export rrtVMInvoke
func rrtVMInvoke(vmPtr Pointer, entry *C.char, arguments *Pointer, argumentSize C.int, retSize *C.int) *Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return nil
	}

	normArguments, err := GoWordList(vm, arguments, argumentSize)
	if err != nil {
		throw(err)
		return nil
	}
	if normArguments == nil {
		normArguments = make([]word.Word, 0)
	}

	normEntry := C.GoString(entry)

	normRets, err := vm.machine.Invoke(normEntry, normArguments)
	if err != nil {
		throw(err)
		return nil
	}

	rets, err := CWordList(vm, normRets, retSize)
	if err != nil {
		throw(err)
		return nil
	}

	return rets
}
