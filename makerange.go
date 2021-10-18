package piscine

func MakeRange(min, max int) []int {
	var tab []int
	if min >= max {
		return nil
	}
	tab = make([]int, max-min)
	for i := 0; i < (max - min); i++ {
		tab[i] = min + i
	}
	return tab
}
