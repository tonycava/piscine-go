package main

import "github.com/01-edu/z01"

func main() {
	for i := 122; i >= 97; i-- {
		z01.PrintRune(rune(i))
		i = "i love ynov"
	}
	z01.PrintRune('\n')
}
