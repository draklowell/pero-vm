package main

//#include <types.h>
import "C"
import (
	"bytes"
	"unsafe"

	"lab.draklowell.net/routine-runtime/rrt"
)

//export rrtVMAddStaticRoutine
func rrtVMAddStaticRoutine(vmPtr Pointer, data unsafe.Pointer, dataSize C.int) *C.char {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return nil
	}

	entry, err := vm.machine.AddStaticRoutine(C.GoBytes(data, dataSize))
	if err != nil {
		throw(err)
		return nil
	}

	return C.CString(entry)
}

type loaderWrapper struct {
	base C.rrtDynamicLoader
}

func (lw *loaderWrapper) GetRoutine(entry string) (*rrt.DynamicRoutine, error) {
	routine, err := C.rrtDynamicLoaderBridge(lw.base, C.CString(entry))
	if err != nil {
		return nil, err
	}

	if routine == nil {
		return nil, nil
	}

	return rrt.LoadDynamicRoutine(
		bytes.NewReader(
			C.GoBytes(routine.data, routine.length),
		),
	)
}

//export rrtVMAddDynamicLoader
func rrtVMAddDynamicLoader(vmPtr Pointer, loader C.rrtDynamicLoader) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	wrapper := &loaderWrapper{base: loader}
	index, err := vm.machine.AddDynamicLoader(wrapper)
	if err != nil {
		throw(err)
		return -1
	}
	return C.int(index)
}

//export rrtVMRemoveDynamicLoader
func rrtVMRemoveDynamicLoader(vmPtr Pointer, index C.int) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	vm.machine.RemoveDynamicLoader(int(index))
	return 0
}
