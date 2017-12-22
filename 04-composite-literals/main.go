package main

import "fmt"

func main() {
	a := [...]string{0: "no error", 1: "Eio", 2: "invalid argument"}
	s := []string{0: "no error", 1: "Eio", 2: "invalid argument"}
	// m := map[str]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}

	fmt.Printf("%T %v\n", a, a)
	fmt.Printf("%T %v\n", s, s)
}
