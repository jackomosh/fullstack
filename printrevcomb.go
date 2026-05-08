package main

import "github.com/01-edu/z01"

func main() {
	
	for i := 9; i >= 2; i-- {
		for j := i - 1; j >= 1; j-- {
			for k := j - 1; k >= 0; k-- {
				if !(i == 9 && j == 8 && k == 7) {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
				z01.PrintRune(rune(i+'0'))
				z01.PrintRune(rune(j+'0'))
				z01.PrintRune(rune(k+'0'))
			}
		}
	}
	z01.PrintRune('\n')
}
