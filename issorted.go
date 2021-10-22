package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	mybool := true
	yourbool := true
	for i := 1; i < len(a); i++ {
			if f(a[i-1], a[i]) >= 0 {
				mybool = false
			} else {
				yourbool = false
			}
		}
	}
	return mybool || yourbool
}
