package main

type MyReader struct{}

func (r MyReader) Read(b []byte) (n int, e error) {
	for i := range b {
		b[i] = byte('A')
	}
	return len(b), nil
}
