package main

import "fmt"

func HashCode(dec string) string {

	size := len(dec)
	res := ""

	for _, ch := range dec {
		hashedValue := int(ch) + size % 127

		if hashedValue < 32 {
			hashedValue += 33
		}

		res += string(rune(hashedValue))
	}
	return res
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}