package main

//#include <types.h>
import "C"
import (
	"bytes"
	"unsafe"

	"lab.draklowell.net/pero-core/pero"
)

//export peroVMAddStaticRoutine
func peroVMAddStaticRoutine(vmPtr Pointer, data unsafe.Pointer, dataSize C.int) *C.char {
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
	base C.peroDynamicLoader
}

func (lw *loaderWrapper) GetRoutine(entry string) (*pero.DynamicRoutine, error) {
	routine, err := C.peroDynamicLoaderBridge(lw.base, C.CString(entry))
	if err != nil {
		return nil, err
	}

	if routine == nil {
		return nil, nil
	}

	return pero.LoadDynamicRoutine(
		bytes.NewReader(
			C.GoBytes(routine.data, routine.length),
		),
	)
}

//export peroVMAddDynamicLoader
func peroVMAddDynamicLoader(vmPtr Pointer, loader C.peroDynamicLoader) C.int {
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

//export peroVMRemoveDynamicLoader
func peroVMRemoveDynamicLoader(vmPtr Pointer, index C.int) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	vm.machine.RemoveDynamicLoader(int(index))
	return 0
}
