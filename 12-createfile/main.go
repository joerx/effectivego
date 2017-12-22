package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := "test.txt"
	if file, err := os.Create(path); err != nil {
		log.Fatalf("Failed to open file %s for writing", path)
	} else {
		fmt.Fprintf(file, "%#v\n", os.Args)
	}
}
