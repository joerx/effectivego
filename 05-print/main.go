package main

import "fmt"

type T struct {
	name  string
	size  int
	value float64
}

func main() {
	t := &T{"john", 12, 3.5}
	s := "someweirdstring"
	var i uint64 = 1<<64 - 1

	fmt.Printf("%+v\n", *t)
	fmt.Printf("% x\n", s)
	fmt.Printf("%x\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", int64(i))
}
