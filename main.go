package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "missing arguments OwO")
	}
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "too many arguments UwU")
	}

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . [STRING]")
		return
	}

	plainTxt, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fileLines := strings.Split(string(plainTxt), "\n")

	// Fetch the input from command-line arguments and clean it
	userInput := cleanstring(os.Args[1])

	if len(userInput) == 0 {
		return
	} else if userInput == "\\n" {
		fmt.Println()
		return
	}
	// Split the input based on the newline delimiter
	inputWords := strings.Split(userInput, "\\n")

	if isempty(inputWords) {
		inputWords = inputWords[:len(inputWords)-1]
	}
	// Iterate through each word and process it
	for _, word := range inputWords {
		if word == "" || word == "\n" {
			fmt.Println()
			continue
		}

		// Render each word line by line
		for lineOffset := 0; lineOffset < 8; lineOffset++ {
			var renderedLine []string

			for _, runeChar := range word {
				startIndex := int((runeChar-32)*9) + 1
				renderedLine = append(renderedLine, fileLines[startIndex+lineOffset])
			}

			// Output the constructed line by joining the slice
			fmt.Println(strings.Join(renderedLine, ""))
		}
	}
}

// cleanstring removes all unprintable characters from a string
func cleanstring(s string) string {
	var cleanString []rune

	for _, r := range s {
		if r >= 32 && r <= 126 { // Check if the character is printable
			cleanString = append(cleanString, r)
		}
	}

	return string(cleanString)
}

func isempty(text []string) bool {
	for _, l := range text {
		if l != "" {
			return false
		}
	}
	return true
}
