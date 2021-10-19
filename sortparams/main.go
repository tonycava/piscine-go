package main

import (
	"github.com/01-edu/z01"
	"os"
)

func sort(table []string) []string {
	var first string
	var second string

	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table)-1; j++ {
			first = table[j]
			second = table[j+1]
			if first > second {
				table[j] = second
				table[j+1] = first
			}
		}
	}
	return table
}

func main() {
	sort(os.Args[1:])
	for _, n := range sort(os.Args[1:]) {
		for i := 0; i < len(n); i++ {
			z01.PrintRune(rune(n[i]))
		}
		z01.PrintRune('\n')
	}
}
