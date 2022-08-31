package word

type Container struct {
	value map[string]Word
}

func NewContainer() *Container {
	return &Container{
		value: make(map[string]Word),
	}
}

func (word *Container) GetType() int {
	return TypeContainer
}

func (word *Container) Get(key string) (Word, error) {
	value, ok := word.value[key]
	if !ok {
		return nil, ErrInvalidKey
	}

	return value, nil
}

func (word *Container) Set(key string, value Word) {
	word.value[key] = value
}

func (word *Container) GetKeys() []string {
	keys := make([]string, 0, len(word.value))
	for key := range word.value {
		keys = append(keys, key)
	}
	return keys
}
