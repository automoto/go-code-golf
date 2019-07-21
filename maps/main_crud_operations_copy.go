package main

import "fmt"

// defining a map
// map[key_type]value_type

func main() {
	//empty init style with "zero" values
	// var colors map[string]string

	//direct declaration and assign style
	colors := map[string]string{
		"red": "#ff0000",
		"green" : "#45b467",
		"black": "#000000",
	}

	fmt.Println(colors)
	//add key value pairs, note the types must match what was declared
	colors["white"] = "#ffffff"

	//remove a key value pair
	delete(colors, "red")
	fmt.Printf("Updated colors: %v \n", colors)

}
