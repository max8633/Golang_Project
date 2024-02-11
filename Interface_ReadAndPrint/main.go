package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// #step1 read a file from Args provied by user
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
