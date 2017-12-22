package main

import (
	"fmt"
)

// all constant have the same value as the first one
const (
	Sunday  = 1 // 1
	Monday      // 1
	Tuesday     // 1
)

// iota is a sequential number generator, incremented by one each time
const (
	_   = iota
	Sun = iota
	Mon = iota
)

// with constant shorthand the value of expression is repeated for each constant in the block,
// with iota we can easily generate sequences
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << 10 * 1
	MB                    // 1 << 10 * 2
	GB                    // 1 << 10 * 3
	TB                    // etc.
)

// See also:
// - https://github.com/golang/go/wiki/Iota
// - https://golang.org/ref/spec#ConstSpec
// - https://golang.org/doc/effective_go.html#constants

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)

	fmt.Println("--")

	fmt.Println(Sun)
	fmt.Println(Mon)

	fmt.Println("--")

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
