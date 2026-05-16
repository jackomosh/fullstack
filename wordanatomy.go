package main

import (
	"fmt"
)

func WordAnatomy(initialWord string, prefixArray []string, suffixArray []string) string {

	matchedPrefix := ""
	matchedSuffix := ""

	for _, p := range prefixArray {
		if hasPrefix(initialWord, p) {
			matchedPrefix = p
			break
		}
	}

	for _, s := range suffixArray {
		if hasSuffix(initialWord, s) {
			matchedSuffix = s
			break
		}
	}
	return "Prefix: " + matchedPrefix + ", " + "Suffix: " + matchedSuffix 
}

func hasPrefix(initialWord, p string) bool {
	if len(p) > len(initialWord) {
		return false
	}
	return initialWord[:len(p)] == p
}


func hasSuffix(initialWord, s string) bool {
	if len(s) > len(initialWord) {
		return false
	}
	return initialWord[len(initialWord)-len(s):] == s
}



func main() {
	// Test 1: Standard match

	word1 := "unbreakable"
	prefixes1 := []string{"re", "pre", "un"}
	suffixes1 := []string{"ing", "able", "ly"}
	fmt.Println("Test 1:", WordAnatomy(word1, prefixes1, suffixes1))

	// Expected: prefix: un, suffix: able

	// Test 2: No matching prefix, one matching suffix

	word2 := "walking"
	prefixes2 := []string{"dis", "mis"}
	suffixes2 := []string{"ed", "ing"}
	fmt.Println("Test 2:", WordAnatomy(word2, prefixes2, suffixes2))

	// Expected: prefix: , suffix: ing

	// Test 3: Multiple possible matches (should take the FIRST one in the array)
	// Even though "dis" and "distrust" both fit, it picks "dis" because it's first.

	word3 := "distrustful"
	prefixes3 := []string{"dis", "distrust"}
	suffixes3 := []string{"ful", "l"}
	fmt.Println("Test 3:", WordAnatomy(word3, prefixes3, suffixes3))

	// Expected: prefix: dis, suffix: ful

	// Test 4: Empty arrays or no matches

	word4 := "banana"
	fmt.Println("Test 4:", WordAnatomy(word4, []string{}, []string{"z"}))

	// Expected: prefix: , suffix: 
}