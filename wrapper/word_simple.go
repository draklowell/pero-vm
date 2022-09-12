package main

import "lab.draklowell.net/pero-core/common/word"

import "C"

//export peroWordIntegerNew
func peroWordIntegerNew(vmPtr Pointer, value C.long) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	integer, err := vm.machine.Heap().NewInteger(int64(value))
	if err != nil {
		throw(err)
		return NullPointer
	}

	integerPtr, err := vm.words.Add(integer)
	if err != nil {
		throw(err)
		return NullPointer
	}
	return integerPtr
}

//export peroWordIntegerGetValue
func peroWordIntegerGetValue(vmPtr Pointer, integerPtr Pointer, result *C.long) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(integerPtr)
	if err != nil {
		throw(err)
		return -1
	}

	integer, ok := value.(*word.Integer)
	if !ok {
		throw(ErrInvalidWordType)
		return -1
	}

	*result = C.long(integer.GetValue())
	return 0
}

//export peroWordFloatNew
func peroWordFloatNew(vmPtr Pointer, value C.double) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	float, err := vm.machine.Heap().NewFloat(float64(value))
	if err != nil {
		throw(err)
		return NullPointer
	}

	floatPtr, err := vm.words.Add(float)
	if err != nil {
		throw(err)
		return -1
	}
	return floatPtr
}

//export peroWordFloatGetValue
func peroWordFloatGetValue(vmPtr Pointer, floatPtr Pointer, result *C.double) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(floatPtr)
	if err != nil {
		throw(err)
		return -1
	}

	float, ok := value.(*word.Float)
	if !ok {
		throw(ErrInvalidWordType)
		return -1
	}

	*result = C.double(float.GetValue())
	return 0
}

//export peroWordBooleanNew
func peroWordBooleanNew(vmPtr Pointer, value C.char) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	valueBool := false
	if value > 0 {
		valueBool = true
	}

	boolean, err := vm.machine.Heap().NewBoolean(valueBool)
	if err != nil {
		throw(err)
		return NullPointer
	}

	booleanPtr, err := vm.words.Add(boolean)
	if err != nil {
		throw(err)
		return NullPointer
	}

	return booleanPtr
}

//export peroWordTrue
func peroWordTrue(vmPtr Pointer) Pointer {
	return peroWordBooleanNew(vmPtr, 1)
}

//export peroWordFalse
func peroWordFalse(vmPtr Pointer) Pointer {
	return peroWordBooleanNew(vmPtr, 0)
}

//export peroWordBooleanGetValue
func peroWordBooleanGetValue(vmPtr Pointer, booleanPtr Pointer, result *C.char) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	value, err := vm.words.Get(booleanPtr)
	if err != nil {
		throw(err)
		return -1
	}

	boolean, ok := value.(*word.Boolean)
	if !ok {
		throw(ErrInvalidWordType)
		return -1
	}

	if boolean.GetValue() {
		*result = 1
	} else {
		*result = 0
	}
	return 0
}

//export peroWordNone
func peroWordNone(vmPtr Pointer) Pointer {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return NullPointer
	}

	nonePtr, err := vm.words.Add(word.None)
	if err != nil {
		throw(err)
		return NullPointer
	}

	return nonePtr
}
