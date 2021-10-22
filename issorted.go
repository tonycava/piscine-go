package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i, nb := range a {
		if f(nb+i, nb+i+1) > 0 {
			return true
		}
	}
	return false
}
