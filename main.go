package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/utils"
)

func main() {
	args, status := utils.CheckArgs()

	if status != "OK" {
		fmt.Fprintln(os.Stderr, "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	plainTxt, err := os.ReadFile("./assets/banners/" + args.BannerName + ".txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err, "\nUsage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	txt := strings.ReplaceAll(string(plainTxt), "\r\n", "\n")
	fileLines := strings.Split(txt, "\n")
	if len(fileLines) != 856 {
		fmt.Fprintln(os.Stderr, "banner file", args.BannerName, "has been modified and is invalid")
		return
	}
	// Fetch the input from command-line arguments and clean it
	userInput := utils.CleanString(args.Text)

	if len(userInput) == 0 {
		return
	} else if userInput == "\\n" {
		fmt.Println()
		return
	}
	// Split the input based on the newline delimiter
	inputWords := strings.Split(userInput, "\\n")

	if utils.IsEmpty(inputWords) {
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
