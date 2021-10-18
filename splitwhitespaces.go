package piscine

func SplitWhiteSpaces(s string) []string {
	var tap []string
	var temp string
	for i := 0; i < len(s); i++ {
		tap = append(tap, string(temp[i]))
	}
	return tap
}
