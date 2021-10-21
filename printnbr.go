package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {

	table := []rune{}
	var first rune
	var second rune
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
			table = append(table, result)
		}
	}
	if n < 0 {
		z01.PrintRune('-')
		for n <= -1 {
			c := n % 10
			result := '0'
			for i := 0; i > c; i-- {
				result++
			}
			n = n / 10
			table = append(table, result)
		}
	}
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
	for x := 0; x < len(table); x++ {
		z01.PrintRune(table[x])
	}
}
