package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("Error: Could not read file %v: %v\n", os.Args[1], err)
		return
	}
	words := strings.Fields(string(content))
	words = applyTags(words)
	text := strings.Join(words, " ")
	text = fixPunctuation(text)
	text = fixQuotes(text)
	text = fixArticles(text)
	// Ensure final newline for clean 'cat' output
	output := strings.TrimSpace(text) + "\n"
    // 2. Capture the error (note the 'err :=' instead of just '_ =')
    err = os.WriteFile(os.Args[2], []byte(output), 0644)
	if err != nil {
		// Using %v for the error object includes the ": open /root/result.txt: permission denied" part
		fmt.Printf("Error: Could not write to file %v: %v\n", os.Args[2], err)
		return 
	}
}

func applyTags(words []string) []string {
	for i := 0; i < len(words); i++ {
		word := words[i]
		// Multi-word tags: (up, n), (low, n), (cap, n)
		if (word == "(up," || word == "(low," || word == "(cap,") && i+1 < len(words) {
			numStr := strings.TrimRight(words[i+1], ")")
			num, _ := strconv.Atoi(numStr)
			start := i - num
			if start < 0 { start = 0 }

			for j := start; j < i; j++ {
				if word == "(up," { words[j] = strings.ToUpper(words[j]) }
				if word == "(low," { words[j] = strings.ToLower(words[j]) }
				if word == "(cap," { words[j] = capitalize(words[j]) }
			}
			words = append(words[:i], words[i+2:]...)
			i--
			continue
		}
		// Single tags
		switch word {
		case "(hex)":
			if i > 0 {
				val, _ := strconv.ParseInt(words[i-1], 16, 64)
				words[i-1] = fmt.Sprintf("%d", val)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(bin)":
			if i > 0 {
				val, _ := strconv.ParseInt(words[i-1], 2, 64)
				words[i-1] = fmt.Sprintf("%d", val)
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(up)":
			if i > 0 {
				words[i-1] = strings.ToUpper(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(low)":
			if i > 0 {
				words[i-1] = strings.ToLower(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		case "(cap)":
			if i > 0 {
				words[i-1] = capitalize(words[i-1])
				words = append(words[:i], words[i+1:]...)
				i--
			}
		}
	}
	return words
}

func capitalize(s string) string {
	if len(s) == 0 { return "" }
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

func fixPunctuation(text string) string {
	// Drag punctuation to the previous word: "sad , because" -> "sad, because"
	reClose := regexp.MustCompile(`\s+([.,!?:;]+)`)
	text = reClose.ReplaceAllString(text, "$1")
	// Ensure space after punctuation, but not before another punctuation
	reOpen := regexp.MustCompile(`([.,!?:;]+)([^.,!?:;\s])`)
	text = reOpen.ReplaceAllString(text, "$1 $2")
	return text
}

func fixQuotes(text string) string {
	// Handle ' quote ' -> 'quote'
	// This regex captures content between two ' marks and trims surrounding spaces
	re := regexp.MustCompile(`'\s+(.*?)\s+'`)
	return re.ReplaceAllString(text, "'$1'")
}

func fixArticles(text string) string {
	words := strings.Fields(text)
	for i := 0; i < len(words)-1; i++ {
		lowerWord := strings.ToLower(words[i])
		if lowerWord == "a" || lowerWord == "an" {
			nextWord := strings.ToLower(words[i+1])
			// Check if next word starts with a vowel or 'h'
			if strings.ContainsAny(string(nextWord[0]), "aeiouh") {
				if words[i][0] == 'A' { words[i] = "An" } else { words[i] = "an" }
			} else {
				if words[i][0] == 'A' { words[i] = "A" } else { words[i] = "a" }
			}
		}
	}
	return strings.Join(words, " ")
}
