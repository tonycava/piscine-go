package piscine

func Capitalize(s string) string {
	result := ""
	maFrance := true
	for _, n := range s {
		if n >= 'a' && n <= 'z' && maFrance {
			n = n - 32
		} else if n >= 'A' && n <= 'Z' && !maFrance {
			n = n + 32
		}
		maFrance = false
		if (n < 'a' || n > 'z') && (n < 'A' || n > 'Z') && (n < '0' || n > '9') {
			maFrance = true
		}
		result = result + string(n)
	}
	return result
}
