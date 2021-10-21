package piscine

func CountIf(f func(string) bool, tab []string) int {
	var nb int
	for _, sale := range tab {
		if f(sale) == true {
			nb++
		}
	}
	return nb
}
