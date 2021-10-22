package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for j := 0; j < len(a); j++ {
		for _, nb := range a {
			if f(nb+j, nb+j+1) > 0 {
				return true
			}
		}
	}
	return false
}
