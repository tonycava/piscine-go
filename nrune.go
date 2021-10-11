package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	if n > len(runes) || n <= 0 { // gestion d'erreur
		return 0
	} else { // dans le cas ou la gestion d'erreur est bonne
		return runes[n-1]
	}
}
