package main

import "fmt"

func main() {
	f := "%T\t%#U\t%32b\n"
	str := "abc日本\x80語"

	for _, char := range str {
		fmt.Printf(f, char, char, char)
	}

	fmt.Println("--")

	c := 1 << 30
	fmt.Printf("%T %b\n", c, c)
}
