package context

import (
	"errors"

	"lab.draklowell.net/pero-core/common/word"
)

func (ctx *Context) read(size int) ([]byte, error) {
	buffer := make([]byte, size)
	if size == 0 {
		return buffer, nil
	}

	if ctx.offset+size > len(ctx.bytecode) {
		return nil, ErrUnexpectedEnd
	}

	for i := 0; i < size; i++ {
		buffer[i] = ctx.bytecode[ctx.offset+i]
	}

	ctx.offset += size

	return buffer, nil
}

func (ctx *Context) seek(offset int) {
	ctx.offset += offset
}

func (ctx *Context) readCommand() (uint8, bool, error) {
	commandBuffer, err := ctx.read(1)
	if err != nil {
		if errors.Is(err, ErrUnexpectedEnd) {
			return 0, true, nil
		}
		return 0, false, err
	}
	return commandBuffer[0], false, nil
}

func (ctx *Context) readU1() (uint8, error) {
	buffer, err := ctx.read(1)
	if err != nil {
		return 0, err
	}
	return uint8(buffer[0]), nil
}

func (ctx *Context) readU2() (uint16, error) {
	buffer, err := ctx.read(2)
	if err != nil {
		return 0, err
	}
	return ctx.order.Uint16(buffer), nil
}

func (ctx *Context) readU4() (uint32, error) {
	buffer, err := ctx.read(4)
	if err != nil {
		return 0, err
	}
	return ctx.order.Uint32(buffer), nil
}

func (ctx *Context) readS4() (int32, error) {
	base, err := ctx.readU4()
	if err != nil {
		return 0, err
	}
	return int32(base), nil
}

func (ctx *Context) getConstantWord(index uint16) (word.Word, error) {
	if int(index) >= len(ctx.constants) {
		return nil, &ErrConstantNotFound{Index: index}
	}
	constant := ctx.constants[index]
	value, ok := constant.(word.Word)
	if !ok {
		return nil, &ErrInvalidConstantType{Index: index}
	}
	return value, nil
}

func (ctx *Context) getConstantString(index uint16) (string, error) {
	constant := ctx.constants[index]
	if constant == nil {
		return "", &ErrConstantNotFound{Index: index}
	}
	value, ok := constant.(string)
	if !ok {
		return "", &ErrInvalidConstantType{Index: index}
	}
	return value, nil
}
