package piscine

func TrimAtoi(s string) int {
	Isnegative := false
	blank := false
	n := 0
	if len(s) < 1 {
		return 0
	}
	for _, char := range s {
		if char < '0' || char > '9' || (char >= 97 && char <= 122) {
			if char == '-' && !blank {
				Isnegative = true
			}
			continue
		}
		n = n*10 + int(char-48)
		blank = true
	}
	if Isnegative {
		return -n
	}
	return n
}
