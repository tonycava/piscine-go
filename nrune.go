package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n > len(runes) || n <= 0 {
		return 0
	} else {
		return runes[n-1]
	}
}
