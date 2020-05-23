package main

import (
	"bufio"
	"io"
)

func main() {
	text := "Lorem Ipsum"
	scanner := bufio.NewScanner(strings.New(text))

}

type LimitReader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}

func (r *Reader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	r.prevRune = -1
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}
