package context

import "lab.draklowell.net/routine-runtime/internal/word"

const (
	defaultStackSize = 1024
	defaultBlockSize = 1024
)

type Stack struct {
	head      *block
	size      uint16
	maxSize   uint16
	blockSize uint16
}

func NewStack(stackSize uint16) *Stack {
	if stackSize == 0 {
		stackSize = defaultStackSize
	}

	stack := &Stack{
		maxSize:   stackSize,
		blockSize: defaultBlockSize,
	}
	stack.pushBlock()
	return stack
}

func (stack *Stack) Clear() error {
	for stack.size > 1 {
		err := stack.popBlock()
		if err != nil {
			return err
		}
	}
	return nil
}

func (stack *Stack) Dump() []word.Word {
	result := make([]word.Word, 0, stack.size*stack.blockSize)
	b := stack.head
	for b.next != nil {
		result = append(result, b.Dump()...)
	}
	result = append(result, b.Dump()...)
	return result
}

func (stack *Stack) Pop() (word.Word, error) {
	value, err := stack.Fetch()
	if err != nil {
		return value, err
	}

	stack.head.index -= 1
	return value, nil
}

func (stack *Stack) Fetch() (word.Word, error) {
	if err := stack.normalize(); err != nil {
		var result word.Word
		return result, err
	}

	value := stack.head.data[stack.head.index]
	return value, nil
}

func (stack *Stack) Push(element word.Word) error {
	if stack.head.index+1 >= len(stack.head.data) {
		if err := stack.pushBlock(); err != nil {
			return err
		}
	}

	stack.head.index++
	stack.head.data[stack.head.index] = element
	return nil
}

func (stack *Stack) normalize() error {
	if stack.head.index < 0 {
		if err := stack.popBlock(); err != nil {
			return err
		}
		return stack.normalize()
	}
	return nil
}

func (stack *Stack) popBlock() error {
	if stack.head.next == nil {
		return ErrStackEmpty
	}

	stack.head = stack.head.next
	stack.size--
	return nil
}

func (stack *Stack) pushBlock() error {
	if stack.size+1 > stack.maxSize {
		return ErrStackTooLarge
	}

	stack.head = &block{
		index: -1,
		next:  stack.head,
		data:  make([]word.Word, stack.blockSize),
	}
	stack.size++
	return nil
}

type block struct {
	index int
	data  []word.Word
	next  *block
}

func (b *block) Dump() []word.Word {
	result := make([]word.Word, 0, len(b.data))
	index := b.index
	for index >= 0 {
		result = append(result, b.data[index])
		index--
	}
	return result
}
