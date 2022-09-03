package main

//#include <types.h>
import "C"

type breakerWrapper struct {
	base C.rrtBreaker
}

func (bw *breakerWrapper) Run() bool {
	if uint8(C.rrtBreakerBridge(bw.base)) > 0 {
		return true
	} else {
		return false
	}
}

//export rrtVMSetBreaker
func rrtVMSetBreaker(vmPtr Pointer, breaker C.rrtBreaker) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	wrapper := breakerWrapper{base: breaker}
	vm.machine.SetBreaker(wrapper.Run)
	return 0
}

//export rrtVMRemoveBreaker
func rrtVMRemoveBreaker(vmPtr Pointer) C.int {
	vm, err := vms.Get(vmPtr)
	if err != nil {
		throw(err)
		return -1
	}

	vm.machine.RemoveBreaker()
	return 0
}
