package pero

import "lab.draklowell.net/pero-core/internal"

type Breaker func() bool

func (vm *VirtualMachine) SetBreaker(breaker Breaker) {
	vm.machine.Breaker = internal.BreakCallback(breaker)
}

func (vm *VirtualMachine) RemoveBreaker() {
	vm.machine.Breaker = internal.EmptyBreaker
}
