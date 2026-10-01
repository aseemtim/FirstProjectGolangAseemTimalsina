package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	for {
		fmt.Println()
		fmt.Println("----------Menu----------")
		fmt.Println("1. Print every Xth line of a file")
		fmt.Println("2. Count bytes and characters of some text")
		fmt.Println("3. Change a single struct inside a function")
		fmt.Println("4. Change a list of structs inside a function")
		fmt.Println("5. Quit")
		fmt.Println("------------------------")

		choice := readUserInput("Enter your choice: ")
		for choice != "1" && choice != "2" && choice != "3" && choice != "4" && choice != "5" {
			fmt.Println(choice, "is not a valid choice. Please enter a number from 1 to 5.")
			choice = readUserInput("Enter your choice: ")
		}

		switch choice {
		case "1":
			filename := readUserInput("Enter the file name: ")
			x := readUserInput("Enter a line number (prints every nth line): ")
			getContentFromFile(filename, x)
		case "2":
			words := readUserInput("Type or paste some text: ")
			printBytesAndChars(words)
		case "3":
			structFunction()
		case "4":
			operateOnListOfStructs()
		case "5":
			fmt.Println("Thanks for using the program!")
			return
		}
	}
}

// prints prompt and returns everything the user typed
func readUserInput(prompt string) string {
	fmt.Print(prompt)
	input, err := reader.ReadString('\n')
	if err != nil && input == "" {
		fmt.Println("\nSome error occurred while reading input. Exiting...")
		os.Exit(0)
	}
	return strings.TrimRight(input, "\r\n")
}

// print every Xth line of the file
func getContentFromFile(filename string, lineNumber string) {
	x, err := strconv.Atoi(strings.TrimSpace(lineNumber))
	if err != nil {
		fmt.Println(lineNumber, "is not a valid line number")
		return
	}

	content := readFile(filename)
	if content == nil {
		fmt.Println("Could not read file", filename)
		return
	}

	if x <= 0 {
		fmt.Println("Error: Line number must be a positive integer.")
		return
	}

	if x > len(content) {
		fmt.Println("Error: Line number", lineNumber, "exceeds total number of lines.")
		return
	}

	for index := x - 1; index < len(content); index += x {
		fmt.Println("Line", index+1, ":", content[index])
	}
}

func readFile(fileName string) []string {
	content, err := os.ReadFile(fileName)

	if err != nil {
		return nil
	}
	return strings.Split(string(content), "\n")
}

// part 2

func getTotalBytes(values string) int {
	return len(values)
}

func getAllChars(word string) []rune {
	return []rune(word)
}

func getTotalChars(values []rune) int {
	return len(values)
}

func printBytesAndChars(value string) {
	totalChars := getTotalChars(getAllChars(value))
	totalBytes := getTotalBytes(value)
	fmt.Println(value, "is", totalBytes, "bytes")
	fmt.Println("and", totalChars, "characters long.")

}

// part 3
func structFunction() {
	type Phone struct {
		Name          string
		Price         int
		HeadphoneJack bool
	}
	phone1 := Phone{"Iphone", 1200, false}
	fmt.Println("Creation of phone1: ", phone1)

	changeIphone := func(phone Phone) {
		phone.Price = 999
		fmt.Println("Changing phone price and printing the struct:", phone)
	}
	changeIphone(phone1)

	fmt.Println("Printing original struct phone after changing its price through a func:", phone1)
}

// part 4
func operateOnListOfStructs() {
	type Phone struct {
		Name          string
		Price         int
		HeadphoneJack bool
	}
	phone1 := Phone{"Iphone", 1200, false}
	phone2 := Phone{"Samsung", 1200, true}

	listOfPhones := []Phone{phone1, phone2}
	fmt.Println("Printing phones after initial creation:")
	for index, phone := range listOfPhones {
		fmt.Println("Phone", index+1, phone)
	}

	changeStructList := func(phones []Phone) {
		phones[1].Name = "One Plus"
		fmt.Println("Printing phones within the function that changed 2nd struct in the list:")
		for index, phone := range phones {
			fmt.Println("Phone", index+1, phone)
		}
	}
	changeStructList(listOfPhones)
	fmt.Println("Printing phones after changeStructList returned:")
	for index, phone := range listOfPhones {
		fmt.Println("Phone", index+1, phone)
	}

}
