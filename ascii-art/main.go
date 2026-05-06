package main

import (
	"fmt"
	"os"
	"strings"
)

// getBanner reads the template and normalizes line endings
func getBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Normalize Windows CRLF to Unix LF to keep our 9-line math consistent
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	return strings.Split(content, "\n"), nil
}

func main() {
	// 1. Basic argument check
	if len(os.Args) < 2 {
		return
	}

	rawInput := os.Args[1]
	if rawInput == "" {
		return
	}

	// 2. Handle literal backslash-n sequences from shell input
	input := strings.ReplaceAll(rawInput, "\\n", "\n")

	// 3. Special Case: Check if input consists ONLY of newlines
	// This handles inputs like "\n" or "\n\n" specifically
	onlyNewlines := true
	for _, char := range input {
		if char != '\n' {
			onlyNewlines = false
			break
		}
	}

	if onlyNewlines {
		for i := 0; i < len(input); i++ {
			fmt.Println()
		}
		return
	}

	// 4. Load the banner data
	bannerLines, err := getBanner("standard.txt")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 5. Split by newline and process each segment
	inputLines := strings.Split(input, "\n")

	for _, word := range inputLines {
		// If the word is empty, it represents a newline character
		if word == "" {
			fmt.Println()
			continue
		}

		// Print the character art row by row (8 rows per char)
		for i := 1; i <= 8; i++ {
			for _, char := range word {
				// Math: Skip (char-32) blocks of 9 lines, then add the current row 'i'
				calcIndex := int(char-32)*9 + i
				
				if calcIndex >= 0 && calcIndex < len(bannerLines) {
					fmt.Print(bannerLines[calcIndex])
				}
			}
			fmt.Println()
		}
	}
}
