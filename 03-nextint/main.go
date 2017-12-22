package main

import "fmt"

func main() {
	bytes := []byte("A rand0m collect1on 0f d1g1ts and charact3rs")
	var x int
	for i := 0; i < len(bytes); {
		fmt.Println(i)
		x, i = nextInt(bytes, i)
		fmt.Println(x, i)
	}
}

func nextInt(b []byte, i int) (int, int) {
	for ; i < len(b) && !isDigit(b[i]); i++ {
		fmt.Printf("%d: %c\n", i, b[i])
	}
	x := 0
	for ; i < len(b) && isDigit(b[i]); i++ {
		x = x*10 + int(b[i]) - '0'
	}
	return x, i
}

func isDigit(b byte) bool {
	return b >= 0x30 && b <= 0x39
}
