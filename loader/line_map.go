package loader

func loadLineMap(reader *Reader) (map[int]int, error) {
	length, err := reader.ReadU2()
	if err != nil {
		return nil, err
	}

	lineMap := make(map[int]int, length)
	for i := 0; i < int(length); i++ {
		value, err := reader.ReadU4()
		if err != nil {
			return nil, err
		}
		lineMap[int(value)] = i + 1
	}
	return lineMap, nil
}
