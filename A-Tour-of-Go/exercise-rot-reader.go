package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 *rot13Reader) Read(b []byte) (n int, e error) {
	n, e = r13.r.Read(b)
	for i := 0; i < len(b); i++ {
		c := b[i]
		if (c >= 'n' && c <= 'z') || (c >= 'N' && c <= 'Z') {
			b[i] -= 13
		} else {
			b[i] += 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
