package piscine

func AppendRange(min, max int) []int {
	var tab []int
	if min > max {
		return tab
	} else {
		for i := 0; i < max; i++ {
			tab = append(tab, min+i)
		}
	}
	return tab
}
