package piscine

func TrimAtoi(s string) int {
	validstr := valid(s)
	art := 0
	tale := len(validstr)
	signe := 1

	if tale == 0 {
		return 0
	}
	if validstr[0] == '-' {
		signe = -1
	}
	for i := 0; i < tale; i++ {
		if i == 0 && validstr[0] == '-' {
			continue
		}
		if i == 0 && validstr[0] == '+' {
			continue
		}
		if validstr[i] > '9' || validstr[i] < '0' && i != 0 {
			return 0
		} else {
			art *= 10
			art += int(validstr[i]) - '0'
		}
	}
	return art * signe
}
func valid(s string) string {
	table := []rune(s)
	tale := len(table)
	var validstr string
	for i := 0; i < tale; i++ {
		if table[i] >= '0' && table[i] <= '9' || table[i] == '+' || table[i] == '-' {
			validstr = validstr + string(table[i])
		}
	}
	return validstr
}
