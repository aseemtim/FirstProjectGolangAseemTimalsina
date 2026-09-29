package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	getContentFromFile("sillyRec.txt", "0")
}

func getContentFromFile(filename string, lineNumber string) {
	index, err := strconv.Atoi(lineNumber)

	if err != nil {
		fmt.Println(lineNumber, "is not a valid line number")
		return
	}
	index = index - 1

	content := readFile(filename)

	if content == nil {
		fmt.Println("Could not read file", filename)
		return
	}

	if index < 0 {
		fmt.Println("Error: Line number must be a positive integer.")
		fmt.Println("Printing the first line: ")
		index = 0
	}

	if index >= len(content) {
		fmt.Println("Error: Line number", lineNumber, "exceeds total number of lines.")
		fmt.Println("Printing the last line: ")
		index = len(content) - 1
	}

	fmt.Println(content[index])
}

func readFile(fileName string) []string {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return nil
	}
	return strings.Split(string(content), "\n")
}
