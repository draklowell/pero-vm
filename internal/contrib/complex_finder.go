package contrib

import (
	"lab.draklowell.net/routine-runtime/common/word"
	"lab.draklowell.net/routine-runtime/internal"
)

type ComplexFinder struct {
	bases []internal.ModuleFinder
}

func NewComplexFinder(bases []internal.ModuleFinder) *ComplexFinder {
	return &ComplexFinder{
		bases: bases,
	}
}

func (finder *ComplexFinder) AddFinder(base internal.ModuleFinder) {
	finder.bases = append(finder.bases, base)
}

func (finder *ComplexFinder) Execute(machine *internal.Machine, entry string, arguments []word.Word) ([]word.Word, error) {
	for _, base := range finder.bases {
		ret, err := base.Execute(machine, entry, arguments)
		if err != nil {
			return nil, err
		}
		if ret != nil {
			return ret, nil
		}
	}
	return nil, nil
}
