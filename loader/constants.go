package loader

import (
	"lab.draklowell.net/routine-runtime/internal/context"
	"lab.draklowell.net/routine-runtime/word"
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
	return nil, &ErrorUnknownConstantTag{Tag: tag}
}

func loadConstantWord(reader *Reader) (word.Word, error) {
	tag, err := reader.ReadU1()
	if err != nil {
		return nil, err
	}

	switch tag {
	case word.TypeInteger:
		return reader.ReadS8()
	case word.TypeFloat:
		return reader.ReadF8()
	case word.TypeBytes:
		length, err := reader.ReadS4()
		if err != nil {
			return nil, err
		}

		return reader.Read(int(length))
	case word.TypeNone:
		return nil, nil
	}
	return nil, &ErrorUnknownWordTag{Tag: tag}
}
