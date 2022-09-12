package main

import "lab.draklowell.net/pero-core/common/word"

import "C"

//export peroWordBytesNew
func peroWordBytesNew(vmPtr Pointer, value *C.char, valueSize C.int) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	cArray := GoArray(value, valueSize)
	array := make([]byte, len(cArray))
	for i, cByte := range cArray {
		array[i] = byte(cByte)
	}

	bytes, err := vm.machine.Heap().NewBytes(array)
	if err != nil {
		throw(err)
		return NullPointer
	}

	bytesPtr, err := vm.words.Add(bytes)
	if err != nil {
		throw(err)
		return NullPointer
	}
	return bytesPtr
}

//export peroWordBytesGetValue
func peroWordBytesGetValue(vmPtr Pointer, bytesPtr Pointer, valueSize *C.int) *C.char {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return nil
	}

	value, err := vm.words.Get(bytesPtr)
	if err != nil {
		throw(err)
		return nil
	}

	bytes, ok := value.(*word.Bytes)
	if !ok {
		throw(ErrInvalidWordType)
		return nil
	}

	array := bytes.GetValue()
	goArray := make([]C.char, len(array))
	for i, char := range array {
		goArray[i] = C.char(char)
	}

	return CArray(goArray, valueSize)
}
