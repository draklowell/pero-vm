package main

import "lab.draklowell.net/routine-runtime/internal/word"

import "C"

//export rrtWordArrayNew
func rrtWordArrayNew(vmPtr Pointer, size C.int) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	array, err := vm.machine.Heap().NewArray(int(size))
	if err != nil {
		throw(err)
		return NullPointer
	}

	arrayPtr, err := vm.words.Add(array)
	if err != nil {
		throw(err)
		return NullPointer
	}

	return arrayPtr
}

//export rrtWordArrayGet
func rrtWordArrayGet(vmPtr Pointer, arrayPtr Pointer, index C.int) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	value, err := vm.words.Get(arrayPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	array, ok := value.(*word.Array)
	if !ok {
		throw(ErrInvalidWordType)
		return NullPointer
	}

	element, err := array.Get(int(index))
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

//export rrtWordArraySet
func rrtWordArraySet(vmPtr Pointer, arrayPtr Pointer, index C.int, elementPtr Pointer) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(arrayPtr)
	if err != nil {
		throw(err)
		return -1
	}

	array, ok := value.(*word.Array)
	if !ok {
		throw(ErrInvalidWordType)
		return -1
	}

	element, err := vm.words.Get(elementPtr)
	if err != nil {
		throw(err)
		return -1
	}

	err = array.Set(int(index), element)
	if err != nil {
		throw(err)
		return -1
	}
	return 0
}

//export rrtWordArrayLength
func rrtWordArrayLength(vmPtr Pointer, arrayPtr Pointer) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(arrayPtr)
	if err != nil {
		throw(err)
		return -1
	}

	array, ok := value.(*word.Array)
	if !ok {
		throw(ErrInvalidWordType)
		return -1
	}

	return C.int(array.GetSize())
}
