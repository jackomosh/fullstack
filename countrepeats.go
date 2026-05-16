package main

import "fmt"

func CountRepeats(s string) string {
	if s == "" {
		return s
	}

	// Convert input string to runes to handle multi-byte characters safely
	runes := []rune(s)
	var result []rune
	counter := 1

	for i := 0; i < len(runes); i++ {
		// Check if next rune is the same
		if i+1 < len(runes) && runes[i] == runes[i+1] {
			counter++
		} else {
			// Append the current character
			result = append(result, runes[i])

			// Append the counter if > 1
			if counter > 1 {
				// Convert integer to runes manually
				var digits []rune
				temp := counter
				for temp > 0 {
					// '0' is 48 in Unicode. Adding the remainder gives the digit.
					digits = append([]rune{rune(temp%10 + '0')}, digits...)
					temp /= 10
				}
				result = append(result, digits...)
			}
			counter = 1
		}
	}

	return string(result)
}

func main() {
	fmt.Println(CountRepeats("ABCABC"))
	fmt.Println(CountRepeats("AAABBC"))
	fmt.Println(CountRepeats("JjjJohhnnnNn"))
	fmt.Println(CountRepeats("     "))
}
