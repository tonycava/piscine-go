package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 97 && s[i] <= 122 || s[i] >= 65 && s[i] <= 90 || s[i] >= 48 && s[i] <= 57) {
			return false
		}
	}
	return true
}
