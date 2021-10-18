package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	for _, n := range argument[0][2:] {
		z01.PrintRune(n)
	}
	z01.PrintRune('\n')
}
