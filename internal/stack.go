package internal

import (
	"lab.draklowell.net/routine-runtime/common/word"
)

const blockSize = 1024

type Stack struct {
	head      *block
	size      uint
	sizeLimit uint
	blockSize uint16
}

type stackLock struct{}

func (sl *stackLock) GetType() int {
	return -1
}

func NewStack(sizeLimit uint) *Stack {
	stack := &Stack{
		blockSize: blockSize,
		sizeLimit: sizeLimit,
	}
	stack.pushBlock()
	return stack
}

func (stack *Stack) Dump() []word.Word {
	result := make([]word.Word, 0, stack.size*uint(stack.blockSize))
	b := stack.head
	for b.next != nil {
		result = append(result, b.Dump()...)
	}
	result = append(result, b.Dump()...)
	return result
}

func (stack *Stack) PushLock() error {
	return stack.push(&stackLock{})
}

func (stack *Stack) PopLock() error {
	for {
		element, err := stack.pop()
		if err != nil {
			return err
		}

		if element == nil {
			return nil
		}

		if element.GetType() == -1 {
			return nil
		}
	}
}

func (stack *Stack) Push(value word.Word) error {
	return stack.push(value)
}

func (stack *Stack) Pop() (word.Word, error) {
	_, err := stack.Fetch()
	if err != nil {
		return nil, err
	}
	return stack.pop()
}

func (stack *Stack) Fetch() (word.Word, error) {
	value, err := stack.fetch()
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, ErrStackEmpty
	}

	if value.GetType() == -1 {
		return nil, ErrStackEmpty
	}
	return value, nil
}

func (stack *Stack) pop() (word.Word, error) {
	value, err := stack.fetch()
	if err != nil {
		return value, err
	}

	stack.head.index -= 1
	return value, nil
}

func (stack *Stack) fetch() (word.Word, error) {
	if err := stack.normalize(); err != nil {
		var result word.Word
		return result, err
	}

	value := stack.head.data[stack.head.index]
	return value, nil
}

func (stack *Stack) push(element word.Word) error {
	if element == nil {
		return ErrNilPointer
	}

	if stack.head.index+1 >= len(stack.head.data) {
		if err := stack.pushBlock(); err != nil {
			return err
		}
	}

	if stack.getSize() > stack.sizeLimit {
		return ErrStackTooLarge
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
	stack.head = &block{
		index: -1,
		next:  stack.head,
		data:  make([]word.Word, stack.blockSize),
	}
	stack.size++
	return nil
}

func (stack *Stack) getSize() uint {
	return uint(stack.size-1)*uint(stack.blockSize) + uint(stack.head.index+1)
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
