package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]
	parsedTodos := getInput(args)
	outputPath := getOutputPath(args)
	output := fmt.Sprintf("%s \n", getTitle(args))
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

func getTitle(i []string) string{
	return i[1]
}

func getOutputPath(i []string) string{
	if len(i) > 2{
		return i[2]
	}
	defaultPath := os.Getenv("TODO_PATH")
	if !(len(defaultPath) > 1){
		panic("File output path not set!")
	}
	return defaultPath
}

func generateFileName() string{
	now := time.Now()
	return fmt.Sprintf("%d_%d_%d_%d_%d_%d.md", now.Day(), now.Month(), now.Year(), now.Second(),
		now.Minute(), now.Hour())
}

func makeDo(todo string) string{
	return fmt.Sprintf("- [ ] %s \n", strings.Trim(todo, " "))
}

func exportMd(todos string, path string) {
	filePath := fmt.Sprintf("%s/%s", path, generateFileName())
	f, err := os.Create(filePath)
	handleError(err)

	writtenBytes, err := f.WriteString(todos)
	handleError(err)
	fmt.Printf("Wrote %d bytes\n", writtenBytes)
	f.Sync()
}

