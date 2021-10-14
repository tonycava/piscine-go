package piscine

func BasicAtoi(s string) int {
	table := []rune(s)
	art := 0
	bb := len(table)
	for i := 0; i < bb; i++ {
		if table[i] > '9' || table[i] < '0' {
			return art
		} else {
			art *= 10
			art += int(table[i]) - '0'
		}
	}
	return art
}
