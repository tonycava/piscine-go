package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n > 0 {
		for n >= 1 {
			c := n % 10
			result := '0'
			for i := 0; i < c; i++ {
				result++
			}
			n = n / 10
			z01.PrintRune(rune(result))
		}
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	for n <= -1 {
		c := n % 10
		result := '0'
		for i := 0; i > c; i-- {
			result++
		}
		n = n / 10
		z01.PrintRune(rune(result))
	}
}
