package piscine

func Atoi(s string) int {
	table := []rune(s)
	art := 0
	tale := len(table)
	signe := 1
	if table[0] == '-' {
		signe = -1
	}
	for i := 0; i <= tale-1; i++ {
		if i == 0 && table[0] == '-' {
			continue
		}
		if i == 0 && table[0] == '+' {
			continue
		}
		if table[i] > '9' || table[i] < '0' && i != 0 {
			return 0
		} else {
			art *= 10
			art += int(table[i]) - '0'
		}
	}
	return art * signe
}
