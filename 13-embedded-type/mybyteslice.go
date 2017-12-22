package main

import (
	"fmt"
	"io"
)

type myReadWriter interface {
	io.Reader
	io.Writer
}

type myByteSlice []byte

func (mb *myByteSlice) Write(data []byte) (int, error) {
	slice := *mb
	l := len(slice)

	if cap(slice) < l+len(data) {
		newSlice := make([]byte, 2*(l+len(data)))
		copy(newSlice, slice)
		slice = newSlice
	}

	slice = slice[:l+len(data)]
	for i, b := range data {
		slice[i+l] = b
	}

	*mb = slice
	return len(data), nil
}

func (mb *myByteSlice) Read(p []byte) (int, error) {
	n := len(p)
	if len(*mb) < n {
		n = len(*mb)
	}
	copy(p, []byte(*mb)[:n])
	return n, nil
}

func main() {
	src := myByteSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dst := make([]byte, 4)
	src.Read(dst)
	fmt.Println(dst)
}
