package piscine

func SplitWhiteSpaces(s string) []string {
	var tap []string
	var temp string
	for i := 0; i < len(s)-1; i++ {
		temp = temp + string(s[i])
		if s[i] == ' ' {
			tap = append(tap, temp)
			temp = ""
		}
		if i == len(s)-1 {
			tap = append(tap, temp)
		}
	}
	return tap
}
