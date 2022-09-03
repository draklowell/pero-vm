package main

import "unsafe"

//#include <types.h>
import "C"

func GetZero[T any]() T {
	var result T
	return result
}

func GoArray[T any](cArray *T, size C.int) []T {
	if cArray == nil {
		return make([]T, 0)
	}
	return (*[1 << 30]T)(unsafe.Pointer(cArray))[:size:size]
}

func CArray[T any](goArray []T, size *C.int) *T {
	*size = C.int(len(goArray))
	if len(goArray) == 0 {
		return nil
	}
	return &goArray[0]
}

type Pointer C.int

var NullPointer Pointer = -1

func NewPointerSet[T any](size int) (PointerSet[T], error) {
	if size < 1 {
		return nil, ErrInvalidPointerSetSize
	}
	return make(PointerSet[T], size), nil
}

type PointerSet[T any] []*T

func (PointerSet PointerSet[T]) Get(ptr Pointer) (T, error) {
	iptr := int(ptr)

	if iptr >= len(PointerSet) || iptr < 0 {
		return GetZero[T](), ErrInvalidPointer
	}

	value := PointerSet[iptr]
	if value == nil {
		return GetZero[T](), ErrInvalidPointer
	}

	return *value, nil
}

func (PointerSet PointerSet[T]) Set(ptr Pointer, value T) error {
	iptr := int(ptr)

	if iptr >= len(PointerSet) || iptr < 0 {
		return ErrInvalidPointer
	}

	PointerSet[iptr] = &value
	return nil
}

func (PointerSet PointerSet[T]) Remove(ptr Pointer) error {
	iptr := int(ptr)

	if iptr >= len(PointerSet) || iptr < 0 {
		return ErrInvalidPointer
	}

	PointerSet[iptr] = nil
	return nil
}

func (PointerSet PointerSet[T]) Add(value T) (Pointer, error) {
	for iptr, element := range PointerSet {
		if element == nil {
			PointerSet[iptr] = &value
			return Pointer(iptr), nil
		}
	}
	return NullPointer, ErrPointerSetLimitExceeded
}

type PointerSetFilter[T any] func(T) bool

func (PointerSet PointerSet[T]) Find(filter PointerSetFilter[T]) Pointer {
	for iptr, element := range PointerSet {
		if filter(*element) {
			return Pointer(iptr)
		}
	}
	return NullPointer
}
