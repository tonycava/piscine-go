package main

import (
	"os"

	"github.com/01-edu/z01"
)

func BasicAtoi(s string) int {
	table := []rune(s)
	art := 0
	bb := len(table)
	for i := 0; i < bb; i++ {
		if table[i] > '9' || table[i] < '0' {
			return 0
		} else {
			art *= 10
			art += int(table[i]) - '0'
		}
	}
	return art
}

func main() {
	arg := os.Args
	IsBlank := false
	upper := 96
	if len(arg) <= 1 {
		return
	}
	if arg[1] == "--upper" {
		upper = 64
	}
	for i := 1; i < len(os.Args); i++ {
		if BasicAtoi(arg[i]) >= 1 && BasicAtoi(arg[i]) <= 26 {
			z01.PrintRune(rune(BasicAtoi(arg[i]) + upper))
			IsBlank = true
		} else if BasicAtoi(arg[i]) == upper || BasicAtoi(arg[i]) == 56 {
			z01.PrintRune(' ')
		}
	}
	if IsBlank == true {
		z01.PrintRune('\n')
	}
}
