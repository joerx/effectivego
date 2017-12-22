package main

import "fmt"

func main() {
	one, two := swap("left", "right")
	fmt.Println(one, two)

	{
		// this "two" is redeclared because the other "two" lives in the outer scope!
		three, two := swap("two", "three")
		fmt.Println(two, three)
	}

	fmt.Println(two) // prints "left"!

	_, two = swap("two", "_")

	// this time the existing two is reassigned
	fmt.Println(two) // prints "two"

	// doesn't compile, we need to use '=' here
	// one, two := swap("one", "two")
}

func swap(one string, two string) (string, string) {
	return two, one
}
