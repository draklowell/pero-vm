package word

func absoluteIndex(index int, size int) (int, error) {
	if index < 0 {
		index = size - index
	}

	if index < 0 || index >= size {
		return 0, ErrInvalidIndex
	}

	return index, nil
}

type Array struct {
	value []Word
}

func NewArray(size int) *Array {
	return &Array{
		value: make([]Word, size),
	}
}

func (word *Array) GetType() int {
	return TypeFloat
}

func (word *Array) Get(index int) (Word, error) {
	index, err := absoluteIndex(index, word.GetSize())
	if err != nil {
		return nil, err
	}

	return word.value[index], nil
}

func (word *Array) Set(index int, value Word) error {
	index, err := absoluteIndex(index, word.GetSize())
	if err != nil {
		return err
	}

	word.value[index] = value
	return nil
}

func (word *Array) GetSize() int {
	return len(word.value)
}
