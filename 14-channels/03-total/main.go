package main

import "fmt"

func total(in chan int, out chan int) {
	total := 0
	for i := range in {
		total += i
	}
	out <- total
}

func main() {
	chIn := make(chan int)
	chOut := make(chan int)
	go total(chIn, chOut)

	chIn <- 1
	chIn <- 2
	chIn <- 3
	close(chIn)

	fmt.Println(<-chOut)
}
