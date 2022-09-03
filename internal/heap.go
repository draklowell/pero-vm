package internal

import (
	"runtime"

	"lab.draklowell.net/routine-runtime/internal/word"
)

type Heap struct {
	sizeLimit uint
	size      uint
}

func NewHeap(sizeLimit uint) *Heap {
	return &Heap{
		sizeLimit: sizeLimit,
	}
}

func (heap *Heap) register(value word.Word) error {
	var size uint = 1
	if bytes, ok := value.(*word.Bytes); ok {
		size = uint(len(bytes.GetValue()))
	}

	if heap.size+size > heap.sizeLimit {
		return ErrHeapTooLarge
	}

	heap.size += size
	runtime.SetFinalizer(value, func(pointer word.Word) {
		heap.size -= size
	})

	return nil
}

func (heap *Heap) NewArray(size int) (*word.Array, error) {
	wordValue := word.NewArray(size)
	if err := heap.register(wordValue); err != nil {
		return nil, err
	}

	return wordValue, nil
}

func (heap *Heap) NewContainer() (*word.Container, error) {
	wordValue := word.NewContainer()
	if err := heap.register(wordValue); err != nil {
		return nil, err
	}

	return wordValue, nil
}

func (heap *Heap) NewInteger(value int64) (*word.Integer, error) {
	wordValue := word.NewInteger(value)
	if err := heap.register(wordValue); err != nil {
		return nil, err
	}

	return wordValue, nil
}

func (heap *Heap) NewFloat(value float64) (*word.Float, error) {
	wordValue := word.NewFloat(value)
	if err := heap.register(wordValue); err != nil {
		return nil, err
	}

	return wordValue, nil
}

func (heap *Heap) NewBytes(value []byte) (*word.Bytes, error) {
	wordValue, err := word.NewBytes(value)
	if err != nil {
		return nil, err
	}

	if err := heap.register(wordValue); err != nil {
		return nil, err
	}

	return wordValue, nil
}

func (heap *Heap) NewBoolean(value bool) (*word.Boolean, error) {
	return word.NewBoolean(value), nil
}
