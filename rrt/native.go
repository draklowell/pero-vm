package rrt

import (
	"lab.draklowell.net/routine-runtime/internal"
	"lab.draklowell.net/routine-runtime/internal/word"
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
