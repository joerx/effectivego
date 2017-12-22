package main

import (
	"fmt"
	"time"
)

// Request represents a request
type Request string

var maxCap = 5
var sem = make(chan int, maxCap)

// Simulate a long operation: wait 2 secs, then print the request and exit
func process(r *Request) {
	t1 := time.NewTimer(time.Second * 2)
	<-t1.C
	fmt.Println(*r)
}

// Serve simulates a server - accept requests from a queue and process them
// only a limited number of requests can be processed at a time, this is controlled by a channel acting as semaphore
// once the channel is drained, it will signal the 'done' channel
func Serve(queue chan *Request, done chan int) {
	for r := range queue {
		sem <- 1
		go func(r *Request) {
			process(r)
			<-sem
		}(r)
	}
	done <- 1
}

func main() {
	queue := make(chan *Request, 20)
	done := make(chan int)

	go Serve(queue, done)

	for i := 0; i < 20; i++ {
		req := Request(fmt.Sprintf("request-%d", i))
		// on an unbuffered channel this will block each time Serve stops reading from the channel
		// on a buffered channel, this will fill up the channel buffer before blocking
		queue <- &req
	}

	fmt.Println("done sending")

	close(queue)
	<-done
}
