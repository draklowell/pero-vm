package rrt

import (
	"lab.draklowell.net/routine-runtime/common/word"
	"lab.draklowell.net/routine-runtime/internal"
)

type NativeRoutine func(*VirtualMachine, []word.Word) ([]word.Word, error)

type nativeRoutineHolder struct {
	routine NativeRoutine
	vm      *VirtualMachine
}

func (nrh *nativeRoutineHolder) Execute(machine *internal.Machine, arguments []word.Word) ([]word.Word, error) {
	return nrh.routine(nrh.vm, arguments)
}

func (vm *VirtualMachine) AddNativeRoutine(entry string, routine NativeRoutine) {
	vm.nativeFinder.SetRoutine(entry, &nativeRoutineHolder{
		routine: routine,
		vm:      vm,
	})
}
