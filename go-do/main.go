package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	parsedTodos := getInput(args)
	outputPath := getOutputPath(args)
	output := ""
	for _, todo := range parsedTodos {
		output += fmt.Sprintf("%s", makeDo(todo))
	}
	fmt.Println(output)
	exportMd(output, outputPath)
}

func handleError(e error) {
	if e!= nil {
		panic(e)
	}
}

func getInput(i []string) []string{
	parsed := strings.Split(i[0], ",")
	return parsed
}

func getOutputPath(i []string) string{
	return i[1]
}

func makeDo(todo string) string{
	return fmt.Sprintf("- [ ] %s \n", strings.Trim(todo, " "))
}

func exportMd(todos string, path string) {
	f, err := os.Create(path)
	handleError(err)

	writtenBytes, err := f.WriteString(todos)
	handleError(err)
	fmt.Printf("Wrote %d bytes\n", writtenBytes)
	f.Sync()
}

