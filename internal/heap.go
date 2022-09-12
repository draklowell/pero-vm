package internal

import (
	"runtime"

	"lab.draklowell.net/pero-core/common/word"
)

type Heap struct {
	sizeLimit uint
	size      uint
	gcMode    int
}

func NewHeap(sizeLimit uint, gcMode int) *Heap {
	return &Heap{
		sizeLimit: sizeLimit,
		gcMode:    gcMode,
	}
}

func calculateSize(value word.Word) uint {
	var size uint = 1
	if bytes, ok := value.(*word.Bytes); ok {
		size = uint(len(bytes.GetValue()))
	}
	return size
}

func (heap *Heap) unregister(value *word.Word) {
	heap.size -= calculateSize(*value)
}

func (heap *Heap) register(value word.Word) error {
	size := calculateSize(value)

	if heap.size+size > heap.sizeLimit {
		if heap.gcMode == GCFrequent || heap.gcMode == GCRare {
			runtime.GC()
			if heap.size+size > heap.sizeLimit {
				return ErrHeapTooLarge
			}
		} else {
			return ErrHeapTooLarge
		}
	}

	heap.size += size
	runtime.SetFinalizer(&value, heap.unregister)

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
