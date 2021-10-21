package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for _, nb := range a {
		for _, nb2 := range a {
			if f(nb, nb2) > 0 {
				return true
			}
		}
	}
	return false
}
