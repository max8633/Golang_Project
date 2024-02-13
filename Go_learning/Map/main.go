package main

import "fmt"

func main() {
	//	#1
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#000000",
		"black": "#ffffff",
	}

	//	#2
	// var colors map[string]string

	//  #3
	// colors := make(map[string]string)
	// colors["red"] = "red"
	// delete(colors, "red")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
