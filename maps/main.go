package main

import "fmt"

// defining a map
// map[key_type]value_type

func main() {
	colors := map[string]string{
		"red": "#ff0000",
		"green" : "#45b467",
		"black": "#000000",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
