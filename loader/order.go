package loader

import (
	"encoding/binary"
)

var (
	BigEndianMagic    = [4]byte{0xEA, 0xDA, 0xCA, 0xBA}
	LittleEndianMagic = [4]byte{0xBA, 0xCA, 0xDA, 0xEA}
)

func detectOrder(reader *Reader) (binary.ByteOrder, error) {
	magic, err := reader.Read(4)
	if err != nil {
		return nil, err
	}

	// Detecting

	var order binary.ByteOrder
	var detectedMagic [4]byte

	switch magic[0] {
	case BigEndianMagic[0]:
		order = binary.BigEndian
		detectedMagic = BigEndianMagic
	case LittleEndianMagic[0]:
		order = binary.LittleEndian
		detectedMagic = LittleEndianMagic
	default:
		return nil, &ErrInvalidMagicNumber{Number: magic}
	}

	// Validating

	for i, char := range magic {
		if char != detectedMagic[i] {
			return nil, &ErrInvalidMagicNumber{Number: magic}
		}
	}
	return order, nil
}
