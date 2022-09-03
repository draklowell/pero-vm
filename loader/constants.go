package loader

import (
	"lab.draklowell.net/routine-runtime/common/word"
	"lab.draklowell.net/routine-runtime/internal/context"
)

func loadConstants(reader *Reader) ([]context.Constant, error) {
	length, err := reader.ReadU2()
	if err != nil {
		return nil, err
	}

	constants := make([]context.Constant, length)

	for i := uint16(0); i < length; i++ {
		constant, err := loadConstant(reader)
		if err != nil {
			return nil, err
		}
		constants[i] = constant
	}

	return constants, nil
}

const (
	ConstantString = 0
	ConstantWord   = 1
)

func loadConstant(reader *Reader) (context.Constant, error) {
	tag, err := reader.ReadU1()
	if err != nil {
		return nil, err
	}

	switch tag {
	case ConstantString:
		length, err := reader.ReadU2()
		if err != nil {
			return nil, err
		}
		return reader.ReadString(int(length))
	case ConstantWord:
		return loadConstantWord(reader)
	}
	return nil, &ErrUnknownConstantTag{Tag: tag}
}

func loadConstantWord(reader *Reader) (word.Word, error) {
	tag, err := reader.ReadU1()
	if err != nil {
		return nil, err
	}

	switch tag {
	case word.TypeInteger:
		value, err := reader.ReadS8()
		if err != nil {
			return nil, err
		}
		return word.NewInteger(value), nil
	case word.TypeFloat:
		value, err := reader.ReadF8()
		if err != nil {
			return nil, err
		}
		return word.NewFloat(value), nil
	case word.TypeBytes:
		length, err := reader.ReadS4()
		if err != nil {
			return nil, err
		}

		value, err := reader.Read(int(length))
		if err != nil {
			return nil, err
		}

		return word.NewBytes(value)
	case word.TypeNone:
		return nil, nil
	}
	return nil, &ErrUnknownWordTag{Tag: tag}
}
