package loader

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

type Reader struct {
	base io.Reader

	Order binary.ByteOrder
}

func NewReader(base io.Reader, order binary.ByteOrder) *Reader {
	return &Reader{
		base:  base,
		Order: order,
	}
}

func (reader *Reader) Read(size int) ([]byte, error) {
	buffer := make([]byte, size)
	if size == 0 {
		return buffer, nil
	}
	n, err := reader.base.Read(buffer)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return nil, &ErrorUnexpectedEOF{Base: err}
		}
		return nil, &ErrorIOError{Base: err}
	}
	if n != size {
		return nil, &ErrorUnexpectedEOF{Base: io.EOF}
	}
	return buffer, nil
}

func (reader *Reader) ReadU1() (uint8, error) {
	buffer, err := reader.Read(1)
	if err != nil {
		return 0, err
	}
	return buffer[0], nil
}

func (reader *Reader) ReadU2() (uint16, error) {
	buffer, err := reader.Read(2)
	if err != nil {
		return 0, err
	}
	return reader.Order.Uint16(buffer), nil
}

func (reader *Reader) ReadU4() (uint32, error) {
	buffer, err := reader.Read(4)
	if err != nil {
		return 0, err
	}
	return reader.Order.Uint32(buffer), nil
}

func (reader *Reader) ReadS4() (int32, error) {
	value, err := reader.ReadU4()
	if err != nil {
		return 0, err
	}
	return int32(value), nil
}

func (reader *Reader) ReadU8() (uint64, error) {
	buffer, err := reader.Read(8)
	if err != nil {
		return 0, err
	}
	return reader.Order.Uint64(buffer), nil
}

func (reader *Reader) ReadS8() (int64, error) {
	value, err := reader.ReadU8()
	if err != nil {
		return 0, err
	}
	return int64(value), nil
}

func (reader *Reader) ReadF8() (float64, error) {
	value, err := reader.ReadU8()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(value), nil
}

func (reader *Reader) ReadString(length int) (string, error) {
	buffer, err := reader.Read(length)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}
