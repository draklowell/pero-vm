package main

//#include <types.h>
import "C"

type breakerWrapper struct {
	base C.peroBreaker
}

func (bw *breakerWrapper) Run() bool {
	if uint8(C.peroBreakerBridge(bw.base)) > 0 {
		return true
	} else {
		return false
	}
}

//export peroVMSetBreaker
func peroVMSetBreaker(vmPtr Pointer, breaker C.peroBreaker) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	wrapper := breakerWrapper{base: breaker}
	vm.machine.SetBreaker(wrapper.Run)
	return 0
}

//export peroVMRemoveBreaker
func peroVMRemoveBreaker(vmPtr Pointer) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	vm.machine.RemoveBreaker()
	return 0
}
