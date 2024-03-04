package main

import "fmt"

func main() {
	firstChan := make(chan string)
	secondChan := make(chan string)

	go func() {
		firstChan <- "first channel win!"
	}()

	go func() {
		secondChan <- "second channel win!"
	}()

	select {
	case msgFromFirstChan := <-firstChan:
		fmt.Println(msgFromFirstChan)
	case msgFromSecondChan := <-secondChan:
		fmt.Println(msgFromSecondChan)
	}
}
