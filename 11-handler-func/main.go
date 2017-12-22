package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// HandlerFunc is a func that can act as http handler
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("call to HandlerFunc.ServeHTTP")
	f(w, r) // a func that implements a method that calls the func itself
}

// ArgServer is a useless http server that only prints its args
func ArgServer(w http.ResponseWriter, r *http.Request) {
	log.Printf("call to ArgServer")
	fmt.Fprintf(w, "%#v\n", os.Args[1:])
}

func main() {
	// because ArgServer has the same sig as HandlerFunc we can use a simple type conversion
	log.Printf("launching server")
	http.Handle("/args", HandlerFunc(ArgServer))
	http.ListenAndServe(":8000", nil)
}
