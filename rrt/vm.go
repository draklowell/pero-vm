package rrt

import (
	"fmt"

	"lab.draklowell.net/routine-runtime/common/word"
	"lab.draklowell.net/routine-runtime/internal"
	"lab.draklowell.net/routine-runtime/internal/context"
	"lab.draklowell.net/routine-runtime/internal/contrib"
)

type VirtualMachine struct {
	name string

	finder  *contrib.ComplexFinder
	machine internal.Machine

	nativeFinder  *contrib.NativeFinder
	staticLoader  *contrib.StaticLoader
	staticFinder  *contrib.ContextFinder
	dynamicLoader *DynamicLoader
	dynamicFinder *contrib.ContextFinder
}

const (
	GCFrequent = internal.GCFrequent
	GCRare     = internal.GCRare
)

func NewVirtualMachine(name string, stackSize uint, heapSize uint, gcMode int) *VirtualMachine {
	nativeFinder := contrib.NewNativeFinder(map[string]contrib.NativeRoutine{})

	staticLoader := &contrib.StaticLoader{
		Routines: map[string]*contrib.StaticRoutine{},
	}
	staticFinder := contrib.NewContextFinder(staticLoader)

	dynamicLoader := NewDynamicLoader()
	dynamicFinder := contrib.NewContextFinder(dynamicLoader)

	finder := contrib.NewComplexFinder([]internal.ModuleFinder{
		nativeFinder,
		staticFinder,
		dynamicFinder,
	})
	return &VirtualMachine{
		name:          name,
		finder:        finder,
		machine:       *internal.NewMachine(finder, stackSize, heapSize, gcMode),
		nativeFinder:  nativeFinder,
		staticLoader:  staticLoader,
		staticFinder:  staticFinder,
		dynamicLoader: dynamicLoader,
		dynamicFinder: dynamicFinder,
	}
}

func (vm *VirtualMachine) Invoke(entry string, arguments []word.Word) ([]word.Word, error) {
	result, err := vm.machine.Execute(fmt.Sprintf("<VM:%s>", vm.name), entry, arguments)
	if err != nil {
		if traceErr, ok := err.(*context.ErrTraceBack); ok {
			return nil, &ErrExecution{Base: traceErr}
		}
		return nil, err
	}
	return result, nil
}

func (vm *VirtualMachine) RemoveRoutine(entry string) {
	vm.nativeFinder.SetRoutine(entry, nil)
	vm.staticLoader.Routines[entry] = nil
}

func (vm *VirtualMachine) Heap() *internal.Heap {
	return vm.machine.Heap
}
