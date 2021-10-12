package piscine

func Index(s string, toFind string) int {
	tableau := []rune(s)
	tableau2 := []rune(toFind)
	for i := 0; i < len(s); i++ {
		for k := 0; k < len(toFind); k++ {
			if tableau[i] == tableau2[k] {
				return i
			}
		}
	}
	return -1
}
