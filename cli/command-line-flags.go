package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("w", "", "a word")
	numPtr := flag.Int("n", 10, "an int")
	boolPtr := flag.Bool("b", false, "a bool")

	flag.Parse()

	fmt.Println("word: ", *wordPtr)
	fmt.Println("int: ", *numPtr)
	fmt.Println("bool: ", *boolPtr)
}
