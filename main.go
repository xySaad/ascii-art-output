package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art/utils"
)

func main() {
	bannerName, err := utils.GetBannerName()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	plainTxt, err := os.ReadFile("./assets/banners/" + bannerName + ".txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err, "\nUsage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}
	fileLines := strings.Split(string(plainTxt), "\n")
	if len(fileLines) != 856 {
		fmt.Fprintln(os.Stderr, "banner file", bannerName, "has been modified and is invalid")
		return
	}
	// Fetch the input from command-line arguments and clean it
	userInput := utils.Cleanstring(os.Args[1])

	if len(userInput) == 0 {
		return
	} else if userInput == "\\n" {
		fmt.Println()
		return
	}
	// Split the input based on the newline delimiter
	inputWords := strings.Split(userInput, "\\n")

	if utils.Isempty(inputWords) {
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
