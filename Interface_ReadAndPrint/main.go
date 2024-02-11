package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// #step1 open a file from Args provied by user
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// #step2 bcz return value f is File type, and it implement Reader interface, so we can pass it to Copy directly
	io.Copy(os.Stdout, f)
}
