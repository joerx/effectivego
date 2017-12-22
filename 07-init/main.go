package main

import "fmt"

var message string
var number int

func init() {
	message = "hello"
}

func init() {
	number = 42
}

func main() {
	fmt.Println(message)
	fmt.Println(number)
}
