package loader

import (
	"encoding/binary"
	"io"

	"lab.draklowell.net/routine-runtime/internal/context"
)

const Version = 1

func LoadRoutine(stream io.Reader) ([]byte, []context.Constant, map[int]int, string, binary.ByteOrder, error) {
	reader := NewReader(stream, nil)

	order, err := detectOrder(reader)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}

	reader.Order = order

	version, err := reader.ReadU2()
	if err != nil {
		return nil, nil, nil, "", nil, err
	}
	if version != Version {
		return nil, nil, nil, "", nil, ErrInvalidVersion
	}

	entry, err := loadEntry(reader)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}
	if entry == "" {
		return nil, nil, nil, "", nil, ErrInvalidEntryLength
	}

	constants, err := loadConstants(reader)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}

	lineMap, err := loadLineMap(reader)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}

	bytecode, err := loadCode(reader)
	if err != nil {
		return nil, nil, nil, "", nil, err
	}

	return bytecode, constants, lineMap, entry, order, nil
}
