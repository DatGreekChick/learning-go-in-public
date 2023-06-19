package main

import (
	"golang.org/x/tour/reader"
)

type MyReader struct{}

func (reader MyReader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		// ASCII character 'A' = 65 https://www.asciitable.com/
		bytes[i] = 65
	}

	return len(bytes), nil
}

func main() {
	reader.Validate(MyReader{})
}
