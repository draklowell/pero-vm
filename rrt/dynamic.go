package rrt

import (
	"encoding/binary"
	"io"

	"lab.draklowell.net/routine-runtime/internal/context"
	"lab.draklowell.net/routine-runtime/loader"
)

type Loader interface {
	GetRoutine(entry string) (*DynamicRoutine, error)
}

type DynamicRoutine struct {
	Order     binary.ByteOrder
	Bytecode  []byte
	Constants []context.Constant
	LineMap   map[int]int
	Entry     string
}

type DynamicLoader struct {
	loaders []Loader
}

func NewDynamicLoader() *DynamicLoader {
	return &DynamicLoader{
		loaders: make([]Loader, 128),
	}
}

func (dl *DynamicLoader) GetRoutine(entry string) (binary.ByteOrder, []byte, []context.Constant, map[int]int, string, error) {
	for _, loader := range dl.loaders {
		if loader == nil {
			continue
		}

		routine, err := loader.GetRoutine(entry)
		if err != nil {
			return nil, nil, nil, nil, "", err
		}

		if routine != nil {
			if routine.Entry != entry {
				return nil, nil, nil, nil, "", ErrorInvalidEntry
			}
			return routine.Order, routine.Bytecode, routine.Constants, routine.LineMap, routine.Entry, nil
		}
	}
	return nil, nil, nil, nil, "", nil
}

func (dl *DynamicLoader) AddLoader(loader Loader) (int, error) {
	for i, value := range dl.loaders {
		if value == nil {
			dl.loaders[i] = loader
			return i, nil
		}
	}
	return 0, ErrorLoaderLimitExceeded
}

func (dl *DynamicLoader) RemoveLoader(index int) {
	dl.loaders[index] = nil
}

func (vm *VirtualMachine) AddDynamicLoader(loader Loader) (int, error) {
	return vm.dynamicLoader.AddLoader(loader)
}

func (vm *VirtualMachine) RemoveDynamicLoader(index int) {
	vm.dynamicLoader.RemoveLoader(index)
}

func LoadDynamicRoutine(reader io.Reader) (*DynamicRoutine, error) {
	bytecode, constants, lineMap, realEntry, order, err := loader.LoadRoutine(reader)
	if err != nil {
		return nil, err
	}

	return &DynamicRoutine{
		Order:     order,
		Bytecode:  bytecode,
		Constants: constants,
		LineMap:   lineMap,
		Entry:     realEntry,
	}, nil
}
