package loader

func loadEntry(reader *Reader) (string, error) {
	length, err := reader.ReadU2()
	if err != nil {
		return "", err
	}
	return reader.ReadString(int(length))
}
