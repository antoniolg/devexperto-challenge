package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(s byte) byte {
	if s >= 'a' && s < 'a' + 13 || s >= 'A' && s < 'A' + 13 {
		return s + 13
	} else if s > 'z' - 13 && s <= 'z' || s > 'Z' - 13 && s <= 'Z' {
		return s - 13
	} else {
		return s
	}
}

func (e *rot13Reader) Read(b []byte) (n int, err error) {
	n, err = e.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

