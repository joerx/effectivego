package main

import "fmt"

// ByteSlice represents a slice of bytes
type ByteSlice []byte

// Append appends data to a ByteSlice
func (p *ByteSlice) Append(data ...byte) {
	slice := *p

	var l = len(slice)
	var nl = l + len(data)

	// reallocate slice if current cap less than needed
	// allocate twice the required length for future growth
	if cap(slice) < nl {
		newSlice := make([]byte, nl*2)
		copy(newSlice, slice)
		slice = newSlice
	}

	// resize slice to match needed length
	slice = slice[0:nl]

	// copy data into the slice
	for i, b := range data {
		slice[l+i] = b
	}

	// reassign to pointer
	*p = slice
}

// Append appends data to a ByteSlice
func (p *ByteSlice) Write(data []byte) (n int, e error) {
	p.Append(data...)
	return len(data), nil
}

func main() {
	s := make([]int, 4, 8)
	s = s[0:6]
	for i := 0; i < 6; i++ {
		s[i] = i
	}
	fmt.Println(s)

	bs := ByteSlice{0x1, 0x2, 0x3}

	fmt.Println(cap(bs)) // 3

	bs.Append(0x4, 0x5, 0x6)

	fmt.Println(cap(bs)) // 12 (was reallocated)
	fmt.Println(bs)

	bs2 := make(ByteSlice, 3, 7)
	copy(bs2, []byte{0x1, 0x2, 0x3})

	fmt.Println(cap(bs2)) // 7

	bs2.Append(0x4, 0x5, 0x6)

	fmt.Println(cap(bs2)) // 7 (not changed)
	fmt.Println(bs2)

	cnt, err := bs2.Write([]byte{0x7, 0x8, 0x9})
	fmt.Println(cnt, err, bs2)

	fmt.Fprintf(&bs2, "Hello %s", "world")
	fmt.Println(bs2)
}
