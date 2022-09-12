package main

import (
	"lab.draklowell.net/pero-core/common/word"
)

import "C"

//export peroWordContainerNew
func peroWordContainerNew(vmPtr Pointer) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	container, err := vm.machine.Heap().NewContainer()
	if err != nil {
		throw(err)
		return NullPointer
	}

	containerPtr, err := vm.words.Add(container)
	if err != nil {
		throw(err)
		return NullPointer
	}
	return containerPtr
}

//export peroWordContainerGet
func peroWordContainerGet(vmPtr Pointer, containerPtr Pointer, key *C.char) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	value, err := vm.words.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	container, ok := value.(*word.Container)
	if !ok {
		throw(ErrInvalidWordType)
		return NullPointer
	}

	element, err := container.Get(C.GoString(key))
	if err != nil {
		throw(err)
		return NullPointer
	}

	elementPtr, err := vm.words.Add(element)
	if err != nil {
		throw(err)
		return NullPointer
	}
	return elementPtr
}

//export peroWordContainerSet
func peroWordContainerSet(vmPtr Pointer, containerPtr Pointer, key *C.char, elementPtr Pointer) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	container, ok := value.(*word.Container)
	if !ok {
		throw(ErrInvalidWordType)
		return -1
	}

	element, err := vm.words.Get(elementPtr)
	if err != nil {
		throw(err)
		return -1
	}

	err = container.Set(C.GoString(key), element)
	if err != nil {
		throw(err)
		return -1
	}

	return 0
}

//export peroWordContainerKeysCount
func peroWordContainerKeysCount(vmPtr Pointer, containerPtr Pointer) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	container, ok := value.(*word.Container)
	if !ok {
		throw(ErrInvalidWordType)
		return -1
	}

	return C.int(len(container.GetKeys()))
}

//export peroWordContainerKey
func peroWordContainerKey(vmPtr Pointer, containerPtr Pointer, index C.int) *C.char {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return nil
	}

	value, err := vm.words.Get(vmPtr)
	if err != nil {
		throw(err)
		return nil
	}

	container, ok := value.(*word.Container)
	if !ok {
		throw(ErrInvalidWordType)
		return nil
	}

	return C.CString(container.GetKeys()[int(index)])
}
