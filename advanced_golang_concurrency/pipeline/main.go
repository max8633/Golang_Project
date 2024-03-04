package main

import "fmt"

func sliceNumsToChan(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func sqNums(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out
}

func main() {
	// input data for processing
	nums := []int{1, 2, 3, 4}

	// stage 1: read data from nums and send into channel for next process
	numsChan := sliceNumsToChan(nums)

	// stage 2: read data from numsChan and do square operation and send into next process
	sqChan := sqNums(numsChan)

	// stage 3: read data from sqChan and print it
	for data := range sqChan {
		fmt.Println(data)
	}
}
