package main

import "fmt"

func main() {
	aSlice := []int{1}
	aBool := true
	anInt := 123
	anInt32 := int32(anInt)
	aRune := 'c'

	printType(aSlice)
	printType(aBool)
	printType(anInt)
	printType(anInt32)
	printType(aRune)
}

func printType(n interface{}) {
	switch n.(type) {
	default:
		fmt.Printf("Unexpected type %T\n", n)
	case bool:
		fmt.Printf("A boolean '%t'\n", n)
	case int, int32:
		fmt.Printf("An integer '%d'\n", n)
	case *bool:
		fmt.Printf("Pointer to boolean %t\n", n)
	}
}
