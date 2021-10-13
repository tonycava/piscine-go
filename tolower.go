package piscine

func ToLower(s string) string {
	patoche := []rune{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			patoche = append(patoche, rune(s[i]+32))
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			patoche = append(patoche, rune(s[i]))
		}
		if s[i] < 'A' || s[i] > 'Z' && s[i] < 'a' || s[i] > 'z' {
			patoche = append(patoche, rune(s[i]))
		}
	}
	return string(patoche)
}
