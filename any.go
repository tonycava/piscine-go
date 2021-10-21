package piscine

func Any(f func(string) bool, a []string) bool {
	for _, sale := range a {
		if f(sale) == true {
			return true
		}
	}
	return false
}
