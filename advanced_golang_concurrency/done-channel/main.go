package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) {
	// inifite for loop with select inside, if get data from done chan then return
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("do work!!")
		}
	}
}

func main() {

	done := make(chan bool)

	go doWork(done)

	time.Sleep(time.Second * 2)

	close(done)

	fmt.Println("it's over~")
}
