package piscine

func SplitWhiteSpaces(s string) []string {
	var temp []string
	for _, w := range s {
		if w == ' ' {
			temp = append(temp, string(w))
		}
	}
	return temp
}
