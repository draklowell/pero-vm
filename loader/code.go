package loader

func loadCode(reader *Reader) ([]byte, error) {
	length, err := reader.ReadS4()
	if err != nil {
		return nil, err
	}
	if length < 0 {
		length = 0
	}

	return reader.Read(int(length))
}
