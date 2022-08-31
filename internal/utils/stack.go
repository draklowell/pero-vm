package utils

const (
	defaultStackSize = 1024
	defaultBlockSize = 1024
)

type Stack[T interface{} | struct{}] struct {
	head      *block[T]
	size      uint16
	maxSize   uint16
	blockSize uint16
}

func NewStack[T interface{} | struct{}](stackSize uint16) *Stack[T] {
	if stackSize == 0 {
		stackSize = defaultStackSize
	}

	stack := &Stack[T]{
		maxSize:   stackSize,
		blockSize: defaultBlockSize,
	}
	stack.pushBlock()
	return stack
}

func (stack *Stack[T]) Clear() error {
	for stack.size > 1 {
		err := stack.popBlock()
		if err != nil {
			return err
		}
	}
	return nil
}

func (stack *Stack[T]) Dump() []T {
	result := make([]T, 0, stack.size*stack.blockSize)
	b := stack.head
	for b.next != nil {
		result = append(result, b.Dump()...)
	}
	result = append(result, b.Dump()...)
	return result
}

func (stack *Stack[T]) Pop() (T, error) {
	value, err := stack.Fetch()
	if err != nil {
		return value, err
	}

	stack.head.index -= 1
	return value, nil
}

func (stack *Stack[T]) Fetch() (T, error) {
	if err := stack.normalize(); err != nil {
		var result T
		return result, err
	}

	value := stack.head.data[stack.head.index]
	return value, nil
}

func (stack *Stack[T]) Push(element T) error {
	if stack.head.index+1 >= len(stack.head.data) {
		if err := stack.pushBlock(); err != nil {
			return err
		}
	}

	stack.head.index++
	stack.head.data[stack.head.index] = element
	return nil
}

func (stack *Stack[T]) normalize() error {
	if stack.head.index < 0 {
		if err := stack.popBlock(); err != nil {
			return err
		}
		return stack.normalize()
	}
	return nil
}

func (stack *Stack[T]) popBlock() error {
	if stack.head.next == nil {
		return ErrStackEmpty
	}

	stack.head = stack.head.next
	stack.size--
	return nil
}

func (stack *Stack[T]) pushBlock() error {
	if stack.size+1 > stack.maxSize {
		return ErrStackTooLarge
	}

	stack.head = &block[T]{
		index: -1,
		next:  stack.head,
		data:  make([]T, stack.blockSize),
	}
	stack.size++
	return nil
}

type block[T interface{} | struct{}] struct {
	index int
	data  []T
	next  *block[T]
}

func (b *block[T]) Dump() []T {
	result := make([]T, 0, len(b.data))
	index := b.index
	for index >= 0 {
		result = append(result, b.data[index])
		index--
	}
	return result
}
