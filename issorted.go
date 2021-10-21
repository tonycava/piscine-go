package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for _, nb := range a {
		if f(nb, nb+1) > 0 {
			return true
		}
	}
	return false
}
