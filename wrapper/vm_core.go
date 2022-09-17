package main

import (
	"lab.draklowell.net/pero-core/common/word"
	"lab.draklowell.net/pero-core/pero"
)

//#include <types.h>
import "C"

type VM struct {
	machine *pero.VirtualMachine
	words   PointerSet[word.Word]
}

var vms, _ = NewPointerSet[*VM](128)

//export peroGetVersion
func peroGetVersion() *C.char {
	return C.CString(pero.Version)
}

//export peroVMNew
func peroVMNew(name *C.char, wordPointerSetSize C.int, stackSize C.long, heapSize C.long, gcMode C.int) Pointer {
	wordsCache, err := NewPointerSet[word.Word](int(wordPointerSetSize))
	if err != nil {
		throw(err)
		return NullPointer
	}

	ptr, err := vms.Add(&VM{
		machine: pero.NewVirtualMachine(
			C.GoString(name),
			uint(stackSize),
			uint(heapSize),
			int(gcMode),
		),
		words: wordsCache,
	})

	if err != nil {
		throw(err)
		return NullPointer
	}
	return ptr
}

//export peroVMClose
func peroVMClose(vmPtr Pointer) C.int {
	err := vms.Remove(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}
	return 0
}

//export peroVMRemoveRoutine
func peroVMRemoveRoutine(vmPtr Pointer, entry *C.char) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	vm.machine.RemoveRoutine(C.GoString(entry))
	return 0
}

//export peroVMInvoke
func peroVMInvoke(vmPtr Pointer, entry *C.char, arguments *Pointer, argumentSize C.int, retSize *C.int) *Pointer {
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
