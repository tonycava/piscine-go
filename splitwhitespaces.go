package piscine

func SplitWhiteSpaces(s string) []string {
	var tap []string
	for i := 0; i < len(s); i++ {
		tap = append(tap, string(s[i]))
	}
	return tap
}
