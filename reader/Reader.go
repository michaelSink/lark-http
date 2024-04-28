package ByteReader

import "io"

type ByteReader struct {
	Buffer   []byte
	Position int
}

func (reader *ByteReader) ReadLine() (string, error) {
	if reader.Position >= len(reader.Buffer) {
		return "", io.EOF
	}

	line := ""
	for index := reader.Position; index < len(reader.Buffer); index++ {
		currentByte := reader.Buffer[index]

		if currentByte == 13 && index+1 < len(reader.Buffer) && reader.Buffer[index+1] == 10 {
			reader.Position += 2
			return line, nil
		}

		line += string(currentByte)
		reader.Position++
	}

	return line, nil
}

func (reader *ByteReader) Buffered() int {
	return len(reader.Buffer) - reader.Position
}

func (reader *ByteReader) Peek(peek int) []byte {
	if reader.Position+peek >= len(reader.Buffer) {
		return reader.Buffer[reader.Position:len(reader.Buffer)]
	}

	return reader.Buffer[reader.Position : reader.Position+peek]
}
