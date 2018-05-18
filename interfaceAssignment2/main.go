package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Incorrect number of arguments. Required: 1 Recieved: %d\n", len(os.Args)-1)
		os.Exit(1)
	}

	oFile, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(2)
	}

}
