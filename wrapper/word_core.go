package main

import (
	"unsafe"

	"lab.draklowell.net/pero-core/common/word"
)

import "C"

func CWordList(vm *VM, list []word.Word, size *C.int) (*Pointer, error) {
	ptrList := make([]Pointer, len(list))
	for i, ret := range list {
		wordPtr, err := vm.words.Add(ret)
		if err != nil {
			return nil, err
		}

		ptrList[i] = wordPtr
	}

	return CArray(ptrList, size), nil
}

func GoWordList(vm *VM, list *Pointer, size C.int) ([]word.Word, error) {
	array := GoArray(list, size)
	result := make([]word.Word, len(array))
	for i, ptr := range array {
		value, err := vm.words.Get(ptr)
		if err != nil {
			return nil, err
		}

		result[i] = value
	}

	return result, nil
}

//export peroWordId
func peroWordId(vmPtr Pointer, wordPtr Pointer) C.long {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(Pointer(wordPtr))
	if err != nil {
		throw(err)
		return -1
	}

	return C.long(uintptr(unsafe.Pointer(&value)))
}

//export peroWordFree
func peroWordFree(vmPtr Pointer, wordPtr Pointer) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	err = vm.words.Remove(Pointer(wordPtr))
	if err != nil {
		throw(err)
		return -1
	}

	return 0
}

//export peroWordDuplicate
func peroWordDuplicate(vmPtr Pointer, wordPtr Pointer) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	value, err := vm.words.Get(wordPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	newPtr, err := vm.words.Add(value)
	if err != nil {
		throw(err)
		return NullPointer
	}

	return newPtr
}

//export peroWordType
func peroWordType(vmPtr Pointer, wordPtr Pointer) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(wordPtr)
	if err != nil {
		throw(err)
		return -1
	}

	return C.int(value.GetType())
}
