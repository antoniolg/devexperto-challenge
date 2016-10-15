package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (e MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return 8, nil
}

func main() {
	reader.Validate(MyReader{})
}
