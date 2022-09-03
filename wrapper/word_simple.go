package main

import "lab.draklowell.net/routine-runtime/common/word"

import "C"

//export rrtWordIntegerNew
func rrtWordIntegerNew(vmPtr Pointer, value C.long) Pointer {
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

//export rrtWordIntegerGetValue
func rrtWordIntegerGetValue(vmPtr Pointer, integerPtr Pointer, result *C.long) C.int {
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

//export rrtWordFloatNew
func rrtWordFloatNew(vmPtr Pointer, value C.double) Pointer {
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

//export rrtWordFloatGetValue
func rrtWordFloatGetValue(vmPtr Pointer, floatPtr Pointer, result *C.double) C.int {
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

//export rrtWordBooleanNew
func rrtWordBooleanNew(vmPtr Pointer, value C.char) Pointer {
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

//export rrtWordTrue
func rrtWordTrue(vmPtr Pointer) Pointer {
	return rrtWordBooleanNew(vmPtr, 1)
}

//export rrtWordFalse
func rrtWordFalse(vmPtr Pointer) Pointer {
	return rrtWordBooleanNew(vmPtr, 0)
}

//export rrtWordBooleanGetValue
func rrtWordBooleanGetValue(vmPtr Pointer, booleanPtr Pointer, result *C.char) C.int {
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

//export rrtWordNone
func rrtWordNone(vmPtr Pointer) Pointer {
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
