package piscine

func AlphaCount(s string) int {
	var count int
	runes := []rune(s)
	for i := 0; i < len(s); i++ {
		if runes[i] >= 97 && runes[i] <= 122 || runes[i] >= 65 && runes[i] <= 90 {
			count++
		}
	}
	return count
}
