package main

import (
	"fmt"
	"net/http"
)

type ctr int

func (c *ctr) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	*c++
	fmt.Fprintf(w, "counter is: %d\n", *c)
}

func main() {
	c := ctr(0)
	http.Handle("/c", &c)
	http.ListenAndServe(":8000", nil)
}
