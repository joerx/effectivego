package main

import (
	"testing"
)

func TestReadSrcLongerThanDest(t *testing.T) {
	l := 4
	src := myByteSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	dst := make([]byte, l)

	cnt, err := src.Read(dst)

	if err != nil {
		t.Error(err)
	}

	if cnt != len(dst) {
		t.Errorf("Expected %d bytes to be read but got %d", len(dst), cnt)
	}

	if len(dst) != l {
		t.Error("Length of dest should not change")
	}

	for i := 0; i < l; i++ {
		if src[i] != dst[i] {
			t.Errorf("Expected %#v at position %d", src[i], i)
		}
	}
}

func TestReadDestLongerThanSrc(t *testing.T) {
	l := 10
	src := make(myByteSlice, l)
	dst := make([]byte, l+20)

	for i := 0; i < len(src); i++ {
		src[i] = byte(i)
	}

	cnt, err := src.Read(dst)

	if err != nil {
		t.Error(err)
	}

	if cnt != len(src) {
		t.Errorf("Expected %d bytes to be read but got %d", len(src), cnt)
	}

	if len(dst) != l+20 {
		t.Error("Length of dest should not change")
	}

	for i := 0; i < l; i++ {
		if src[i] != dst[i] {
			t.Errorf("Expected %#v at position %d", src[i], i)
		}
	}
}

func TestWriteWithoutReallocation(t *testing.T) {
	l := 10    // src len
	c := l * 2 // dst cap

	src := make([]byte, l)
	dst := make(myByteSlice, 0, c)

	for i := 0; i < len(src); i++ {
		src[i] = byte(i)
	}

	cnt, err := dst.Write(src)

	if err != nil {
		t.Error(err)
	}

	if cnt != len(dst) {
		t.Errorf("Expected %d bytes to be written, but was %d", len(dst), cnt)
	}

	if len(dst) != len(src) {
		t.Errorf("Expected dest length to match src length of %d after write ", len(src))
	}

	if cap(dst) != c {
		t.Errorf("Expected dest capacity not to change, but was %d after write", cap(dst))
	}

	for i := 0; i < len(src); i++ {
		if dst[i] != src[i] {
			t.Errorf("Expected %#v at position %d but was %#v", src[i], i, dst[i])
		}
	}
}

func TestWriteWithReallocation(t *testing.T) {
	c := 10    // dest cap
	l := c * 2 // src len

	src := make([]byte, l)
	dst := make(myByteSlice, 0, c)

	for i := 0; i < len(src); i++ {
		src[i] = byte(2 * i)
	}

	cnt, err := dst.Write(src)

	if err != nil {
		t.Error(err)
	}

	if cnt != len(src) {
		t.Errorf("Expected %d bytes to be written, but was %d", len(src), cnt)
	}

	if len(dst) != len(src) {
		t.Errorf("Expected dst len to match src len of %d after write", len(src))
	}

	if cap(dst) == c {
		t.Error("Expected dst cap to have been reallocated")
	}

	for i := 0; i < len(src); i++ {
		if dst[i] != src[i] {
			t.Errorf("Expected %#v at position %d but was %#v", src[i], i, dst[i])
		}
	}
}
