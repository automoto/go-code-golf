package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	parsed := getInput(args)
	fmt.Println(parsed[2])
}

func getInput(i []string) []string{
	parsed := strings.Split(i[0], ",")
	return parsed
}

