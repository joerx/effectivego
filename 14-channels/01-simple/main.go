package main

import (
	"fmt"
	"sort"
	"time"
)

// does nothing for 2 seconds, then terminates
func doSomething() {
	t1 := time.NewTimer(time.Second * 2)
	<-t1.C
}

func main() {
	lst := []int{12, 34, 56, 0, 9, 13, 23, 15}

	c := make(chan int)
	go func(lst []int) {
		doSomething()
		sort.IntSlice(lst).Sort()
		c <- 1
	}(lst)

	// block until func has completed
	// if we skip this, following statement will print an unsorted list!
	<-c

	fmt.Println(lst)
}
