package main

import "fmt"

type myString string

type myThing struct {
	name string
}

func (t myThing) String() string {
	return t.name
}

func main() {
	printMe(1)
	printMe(int32(2))
	printMe("a nimble string")
	printMe(true)
	printMe(myThing{"john smith"}) // create struct of type mything, use {}
	printMe(myString("my string")) // convert string to myString, use ()

	fmt.Println(isMyString("some string"))         // false
	fmt.Println(isMyString(myString("my string"))) // true
}

func isMyString(thing interface{}) bool {
	if _, ok := thing.(myString); ok {
		return true
	}
	return false
}

func printMe(thing interface{}) {
	switch t := thing.(type) {
	case int, int32, int64:
		fmt.Printf("Integer thing %T: %d\n", t, t)
	case string:
		fmt.Printf("String thing %T: %s\n", t, t)
	case bool:
		fmt.Printf("Boolean thing %T: %t\n", t, t)
	case fmt.Stringer:
		fmt.Printf("Stringer thing: %T: %s\n", t, t.String())
	default:
		fmt.Printf("Unknown thing: %#v\n", t)
	}
}
