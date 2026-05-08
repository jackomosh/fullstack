package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
	}

	var res []rune

	absN := n
	if n < 0 {
		absN = -n
	}

	for absN > 0 {
		digit := absN % 10
		res = append(res, rune(digit+'0'))
		absN /= 10
	}

	if isNegative {
		res = append(res, '-')
	}

	// Reverse the slice
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

func main() {
    fmt.Println(Itoa(12345))
    fmt.Println(Itoa(0))
    fmt.Println(Itoa(-1234))
    fmt.Println(Itoa(987654321))
}