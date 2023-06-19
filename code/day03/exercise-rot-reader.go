package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(bytes []byte) (ascii int, err error) {
	ascii, err = reader.r.Read(bytes)
	if err != nil {
		return
	}

	for i := range bytes {
		// take all capital and lowercase letters in the English alphabet
		if bytes[i] >= 'A' && bytes[i] <= 'z' {
			// rotate the ASCII character by 13 places
			bytes[i] += 13

			// if the ASCII character exceeds the range, start from the beginning
			if bytes[i] > 'z' {
				bytes[i] -= 26
			}
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
