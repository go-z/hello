package main

import (
	"io"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(b []byte) (int, error) {
	n, e := rr.r.Read(b)
	for i := 0; i < n; i++ {
		switch {
		case (b[i] > 64 && b[i] < 78) || (b[i] > 96 && b[i] < 110):
			b[i] += 13
		case (b[i] > 77 && b[i] < 91) || (b[i] > 109 && b[i] < 123):
			b[i] -= 13
		}
	}
	return n, e
}
