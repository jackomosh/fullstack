package ascii

import (
	// "fmt"
	"os"
	"strings"
)

// GetBanner reads the template and normalizes line endings
func GetBanner(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	return strings.Split(content, "\n"), nil
}

// GenerateAscii takes the input string and banner name and returns the art as a string
func GenerateAscii(input string, banner string) (string, error) {
	// 1. Load the banner
	bannerLines, err := GetBanner(banner + ".txt")
	if err != nil {
		return "", err
	}

	// 2. Handle literal backslash-n sequences
	input = strings.ReplaceAll(input, "\\n", "\n")

	// 3. Process lines
	inputLines := strings.Split(input, "\n")
	var result strings.Builder

	// Check if input is only newlines
	onlyNewlines := true
	for _, char := range input {
		if char != '\n' {
			onlyNewlines = false
			break
		}
	}
	if onlyNewlines && len(input) > 0 {
		return strings.Repeat("\n", len(input)), nil
	}

	for _, word := range inputLines {
		if word == "" {
			result.WriteString("\n")
			continue
		}

		for i := 1; i <= 8; i++ {
			for _, char := range word {
				calcIndex := int(char-32)*9 + i
				if calcIndex >= 0 && calcIndex < len(bannerLines) {
					result.WriteString(bannerLines[calcIndex])
				}
			}
			result.WriteString("\n")
		}
	}
	return result.String(), nil
}