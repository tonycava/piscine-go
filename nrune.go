package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n >= len(s) {
		return 0
	}
	if n < 0 {
		return runes[len(runes)-1]
	}
	return runes[n]
}
