package contrib

import (
	"lab.draklowell.net/routine-runtime/common/word"
	"lab.draklowell.net/routine-runtime/internal"
)

type NativeRoutine interface {
	Execute(*internal.Machine, []word.Word) ([]word.Word, error)
}

type NativeFinder struct {
	routines map[string]NativeRoutine
}

func NewNativeFinder(routines map[string]NativeRoutine) *NativeFinder {
	return &NativeFinder{
		routines: routines,
	}
}

func (finder *NativeFinder) SetRoutine(entry string, routine NativeRoutine) {
	finder.routines[entry] = routine
}

func (finder *NativeFinder) Execute(machine *internal.Machine, entry string, arguments []word.Word) ([]word.Word, error) {
	routine := finder.routines[entry]
	if routine == nil {
		return nil, nil
	}
	ret, err := routine.Execute(machine, arguments)
	if err != nil {
		return nil, err
	}
	if ret == nil {
		ret = make([]word.Word, 0)
	}
	return ret, nil
}
