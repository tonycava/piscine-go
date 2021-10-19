package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, str := range a {
		for i := 0; i < len(str); i++ {
			z01.PrintRune(rune(str[i]))
		}
		z01.PrintRune('\n')
	}
}
